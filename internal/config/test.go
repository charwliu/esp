package config

import (
	"flag"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"

	"go.vixal.xyz/esp/pkg/capture"
	"go.vixal.xyz/esp/pkg/fs"
	"go.vixal.xyz/esp/pkg/rnd"
)

// define constants used for testing the config package
const (
	TestDataZip  = "/tmp/esp/testdata.zip"
	TestDataURL  = "https://dl.vixal.xyz/qa/testdata.zip"
	TestDataHash = "be394d5bee8a5634d415e9e0663eef20b5604510" // sha1sum
)

var testConfig *Config
var testConfigOnce sync.Once
var testConfigMutex sync.Mutex

func testDataPath(assetsPath string) string {
	return assetsPath + "/testdata"
}

// NewTestOptions inits valid options used for testing
func NewTestOptions() *Options {
	assetsPath := fs.Abs("../../assets")
	storagePath := fs.Abs("../../storage")
	testDataPath := filepath.Join(storagePath, "testdata")

	dbDriver := os.Getenv("PHOTOPRISM_TEST_DRIVER")
	dbDsn := os.Getenv("PHOTOPRISM_TEST_DSN")

	// Config example for MySQL / MariaDB:
	//   dbDriver = MySQL,
	//   dbDsn = "esp:esp@tcp(esp-db:4001)/esp?parseTime=true",

	if dbDriver == "test" || dbDriver == "sqlite" || dbDriver == "" || dbDsn == "" {
		dbDriver = SQLite
		dbDsn = ".test.db"
	}

	c := &Options{
		Name:           "GoshESP",
		Version:        "0.0.0",
		Copyright:      "(c) 2018-2021 Michael Mayer",
		Debug:          true,
		Public:         true,
		ReadOnly:       false,
		AssetsPath:     assetsPath,
		AutoIndex:      -1,
		AutoImport:     7200,
		StoragePath:    testDataPath,
		CachePath:      testDataPath + "/cache",
		TempPath:       testDataPath + "/temp",
		ConfigPath:     testDataPath + "/config",
		DatabaseDriver: dbDriver,
		DatabaseDsn:    dbDsn,
		AdminPassword:  "esp",
	}

	return c
}

// NewTestOptionsError inits invalid options used for testing
func NewTestOptionsError() *Options {
	assetsPath := fs.Abs("../..")
	testDataPath := fs.Abs("../../storage/testdata")

	c := &Options{
		AssetsPath:     assetsPath,
		StoragePath:    testDataPath,
		CachePath:      testDataPath + "/cache",
		TempPath:       testDataPath + "/temp",
		DatabaseDriver: SQLite,
		DatabaseDsn:    ".test-error.db",
	}

	return c
}

func SetNewTestConfig() {
	testConfig = NewTestConfig()
}

// TestConfig inits the global testConfig if it was not already initialised
func TestConfig() *Config {
	testConfigOnce.Do(SetNewTestConfig)

	return testConfig
}

// NewTestConfig inits valid config used for testing
func NewTestConfig() *Config {
	log := L()
	defer log.Debug(capture.Time(time.Now(), "config: new test config created"))

	testConfigMutex.Lock()
	defer testConfigMutex.Unlock()

	c := &Config{
		options: NewTestOptions(),
		token:   rnd.Token(8),
	}

	s := NewSettings()

	if err := os.MkdirAll(c.ConfigPath(), os.ModePerm); err != nil {
		log.Fatal("config:", zap.Error(err))
	}

	if err := s.Save(filepath.Join(c.ConfigPath(), "settings.yml")); err != nil {
		log.Fatal("config:", zap.Error(err))
	}

	if err := c.Init(); err != nil {
		log.Fatal("config:", zap.Error(err))
	}

	c.InitTestDB()

	return c
}

// NewTestErrorConfig inits invalid config used for testing
func NewTestErrorConfig() *Config {
	c := &Config{options: NewTestOptionsError()}

	return c
}

// CliTestContext returns example cli config for testing
func CliTestContext() *cli.Context {
	config := NewTestOptions()

	globalSet := flag.NewFlagSet("test", 0)
	globalSet.Bool("debug", false, "doc")
	globalSet.String("storage-path", config.StoragePath, "doc")
	globalSet.String("backup-path", config.StoragePath, "doc")
	globalSet.String("config-file", config.ConfigFile, "doc")
	globalSet.String("assets-path", config.AssetsPath, "doc")
	globalSet.Int("auto-index", config.AutoIndex, "doc")
	globalSet.Int("auto-import", config.AutoImport, "doc")

	app := cli.NewApp()
	app.Version = "0.0.0"

	c := cli.NewContext(app, globalSet, nil)

	LogError(c.Set("storage-path", config.StoragePath))
	LogError(c.Set("backup-path", config.BackupPath))
	LogError(c.Set("config-file", config.ConfigFile))
	LogError(c.Set("assets-path", config.AssetsPath))
	LogError(c.Set("temp-path", config.TempPath))
	LogError(c.Set("cache-path", config.CachePath))
	LogError(c.Set("admin-password", config.AdminPassword))
	LogError(c.Set("detect-nsfw", "true"))
	LogError(c.Set("auto-index", strconv.Itoa(config.AutoIndex)))
	LogError(c.Set("auto-import", strconv.Itoa(config.AutoImport)))

	return c
}

// RemoveTestData deletes files in import, export, originals and cache folders
func (c *Config) RemoveTestData(t *testing.T) {

	if err := os.RemoveAll(c.TempPath()); err != nil {
		t.Fatal(err)
	}

	if err := os.RemoveAll(c.CachePath()); err != nil {
		t.Fatal(err)
	}
}

// DownloadTestData downloads test data from esp.org server
func (c *Config) DownloadTestData(t *testing.T) {
	if fs.FileExists(TestDataZip) {
		hash := fs.Hash(TestDataZip)

		if hash != TestDataHash {
			if err := os.Remove(TestDataZip); err != nil {
				t.Fatalf("config: %s", err.Error())
			}

			t.Logf("config: removed outdated test data zip file (fingerprint %s)", hash)
		}
	}

	if !fs.FileExists(TestDataZip) {
		t.Logf("config: downloading latest test data zip file from %s", TestDataURL)

		if err := fs.Download(TestDataZip, TestDataURL); err != nil {
			t.Fatalf("config: test data download failed: %s", err.Error())
		}
	}
}

// UnzipTestData in default test folder
func (c *Config) UnzipTestData(t *testing.T) {
	if _, err := fs.Unzip(TestDataZip, c.StoragePath()); err != nil {
		t.Fatalf("config: could not unzip test data: %s", err.Error())
	}
}

// InitializeTestData using testing constant
func (c *Config) InitializeTestData(t *testing.T) {
	defer t.Logf(capture.Time(time.Now(), "config: initialized test data"))

	c.RemoveTestData(t)

	c.DownloadTestData(t)

	c.UnzipTestData(t)
}
