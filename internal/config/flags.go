package config

import (
	"github.com/klauspost/cpuid/v2"
	"github.com/urfave/cli/v2"
)

// GlobalFlags Gosh ESP command-line parameters and flags.
var GlobalFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "debug",
		Usage:   "run in debug mode, shows additional log messages",
		EnvVars: []string{"GOSH_DEBUG"},
	},
	&cli.BoolFlag{
		Name:   "test",
		Hidden: true,
		Usage:  "run in test mode",
	},
	&cli.BoolFlag{
		Name:    "demo",
		Hidden:  true,
		Usage:   "run in demo mode",
		EnvVars: []string{"GOSH_DEMO"},
	},
	&cli.BoolFlag{
		Name:    "sponsor",
		Hidden:  true,
		Usage:   "your continuous support helps to pay for development and operating expenses",
		EnvVars: []string{"GOSH_SPONSOR"},
	},
	&cli.BoolFlag{
		Name:    "public, p",
		Aliases: []string{"p"},
		Usage:   "no authentication required, disables password protection",
		EnvVars: []string{"GOSH_PUBLIC"},
	},
	&cli.BoolFlag{
		Name:    "read-only",
		Aliases: []string{"r"},
		Usage:   "don't modify originals folder; disables import, upload, and delete",
		EnvVars: []string{"GOSH_READONLY"},
	},
	&cli.BoolFlag{
		Name:    "experimental",
		Aliases: []string{"e"},
		Usage:   "enable experimental features",
		EnvVars: []string{"GOSH_EXPERIMENTAL"},
	},
	&cli.StringFlag{
		Name:    "admin-password",
		Usage:   "initial admin `PASSWORD`, min 4 characters",
		EnvVars: []string{"GOSH_ADMIN_PASSWORD"},
	},
	&cli.StringFlag{
		Name:    "config-file, c",
		Usage:   "load initial config options from `FILENAME`",
		EnvVars: []string{"GOSH_CONFIG_FILE"},
	},
	&cli.StringFlag{
		Name:    "config-path",
		Usage:   "config `PATH` containing application settings",
		EnvVars: []string{"GOSH_CONFIG_PATH"},
	},
	&cli.StringFlag{
		Name:    "storage-path",
		Usage:   "storage `PATH` for cache, database files",
		EnvVars: []string{"GOSH_STORAGE_PATH"},
	},
	&cli.StringFlag{
		Name:    "cache-path",
		Usage:   "cache storage `PATH` for sessions and thumbnails",
		EnvVars: []string{"GOSH_CACHE_PATH"},
	},
	&cli.StringFlag{
		Name:    "temp-path",
		Usage:   "temporary `PATH` for storing uploads and downloads",
		EnvVars: []string{"GOSH_TEMP_PATH"},
	},
	&cli.StringFlag{
		Name:    "backup-path",
		Usage:   "backup storage `PATH`",
		EnvVars: []string{"GOSH_BACKUP_PATH"},
	},
	&cli.StringFlag{
		Name:    "assets-path",
		Usage:   "assets `PATH` for static resources like models and templates",
		EnvVars: []string{"GOSH_ASSETS_PATH"},
	},
	&cli.IntFlag{
		Name:    "workers",
		Aliases: []string{"w"},
		Usage:   "adjusts `MAX` number of indexing workers",
		EnvVars: []string{"GOSH_WORKERS"},
		Value:   cpuid.CPU.PhysicalCores / 2,
	},
	&cli.IntFlag{
		Name:    "wakeup-interval",
		Usage:   "background worker wakeup interval in `SECONDS`",
		EnvVars: []string{"GOSH_WAKEUP_INTERVAL"},
	},
	&cli.IntFlag{
		Name:    "auto-index",
		Usage:   "auto indexing safety delay in `SECONDS` (WebDAV)",
		EnvVars: []string{"GOSH_AUTO_INDEX"},
	},
	&cli.IntFlag{
		Name:    "auto-import",
		Usage:   "auto importing safety delay in `SECONDS` (WebDAV)",
		EnvVars: []string{"GOSH_AUTO_IMPORT"},
	},
	&cli.BoolFlag{
		Name:    "disable-backups",
		Usage:   "don't backup photo and album metadata to YAML files",
		EnvVars: []string{"GOSH_DISABLE_BACKUPS"},
	},
	&cli.BoolFlag{
		Name:    "disable-webdav",
		Usage:   "disables built-in WebDAV server",
		EnvVars: []string{"GOSH_DISABLE_WEBDAV"},
	},
	&cli.BoolFlag{
		Name:    "disable-settings",
		Usage:   "users can not view or change settings",
		EnvVars: []string{"GOSH_DISABLE_SETTINGS"},
	},
	&cli.BoolFlag{
		Name:    "disable-places",
		Usage:   "disables reverse geocoding and maps",
		EnvVars: []string{"GOSH_DISABLE_PLACES"},
	},
	&cli.StringFlag{
		Name:    "log-level, l",
		Aliases: []string{"l"},
		Usage:   "trace, debug, info, warning, error, fatal or panic",
		Value:   "info",
		EnvVars: []string{"GOSH_LOG_LEVEL"},
	},
	&cli.StringFlag{
		Name:    "log-filename",
		Usage:   "server log `FILENAME`",
		EnvVars: []string{"GOSH_LOG_FILENAME"},
		Value:   "",
	},
	&cli.StringFlag{
		Name:    "pid-filename",
		Usage:   "server process id `FILENAME`",
		EnvVars: []string{"GOSH_PID_FILENAME"},
	},
	&cli.StringFlag{
		Name:    "cdn-url",
		Usage:   "content delivery network `URL` (optional)",
		EnvVars: []string{"GOSH_CDN_URL"},
	},
	&cli.StringFlag{
		Name:    "site-url",
		Usage:   "public site `URL`",
		Value:   "http://localhost:2342/",
		EnvVars: []string{"GOSH_SITE_URL"},
	},
	&cli.StringFlag{
		Name:    "site-preview",
		Usage:   "public preview image `URL`",
		EnvVars: []string{"GOSH_SITE_PREVIEW"},
	},
	&cli.StringFlag{
		Name:    "site-title",
		Usage:   "site title",
		Value:   "Gosh ESP",
		EnvVars: []string{"GOSH_SITE_TITLE"},
	},
	&cli.StringFlag{
		Name:    "site-caption",
		Usage:   "short site caption",
		Value:   "Browse Your Life",
		EnvVars: []string{"GOSH_SITE_CAPTION"},
	},
	&cli.StringFlag{
		Name:    "site-description",
		Usage:   "long site description",
		EnvVars: []string{"GOSH_SITE_DESCRIPTION"},
	},
	&cli.StringFlag{
		Name:    "site-author",
		Usage:   "site artist or copyright",
		EnvVars: []string{"GOSH_SITE_AUTHOR"},
	},
	&cli.IntFlag{
		Name:    "http-port",
		Value:   2342,
		Usage:   "http server port `NUMBER`",
		EnvVars: []string{"GOSH_HTTP_PORT"},
	},
	&cli.StringFlag{
		Name:    "http-host",
		Usage:   "http server `IP` address",
		EnvVars: []string{"GOSH_HTTP_HOST"},
	},
	&cli.StringFlag{
		Name:    "http-mode",
		Aliases: []string{"m"},
		Usage:   "debug, release or test",
		EnvVars: []string{"GOSH_HTTP_MODE"},
	},
	&cli.StringFlag{
		Name:    "http-compression",
		Aliases: []string{"z"},
		Usage:   "improves transfer speed and bandwidth utilization (none or gzip)",
		EnvVars: []string{"GOSH_HTTP_COMPRESSION"},
	},
	&cli.StringFlag{
		Name:    "database-driver",
		Usage:   "database driver `NAME` (sqlite or mysql)",
		Value:   "sqlite",
		EnvVars: []string{"GOSH_DATABASE_DRIVER"},
	},
	&cli.StringFlag{
		Name:    "database-dsn",
		Usage:   "sqlite file name, specifying a `DSN` is optional for mariadb and mysql",
		EnvVars: []string{"GOSH_DATABASE_DSN"},
	},
	&cli.StringFlag{
		Name:    "database-server",
		Usage:   "database server `HOST`, specifying a :port is optional",
		EnvVars: []string{"GOSH_DATABASE_SERVER"},
	},
	&cli.StringFlag{
		Name:    "database-name",
		Value:   "esp",
		Usage:   "database schema `NAME`",
		EnvVars: []string{"GOSH_DATABASE_NAME"},
	},
	&cli.StringFlag{
		Name:    "database-user",
		Value:   "esp",
		Usage:   "database user `NAME`",
		EnvVars: []string{"GOSH_DATABASE_USER"},
	},
	&cli.StringFlag{
		Name:    "database-password",
		Usage:   "database user `PASSWORD`",
		EnvVars: []string{"GOSH_DATABASE_PASSWORD"},
	},
	&cli.IntFlag{
		Name:    "database-conns",
		Usage:   "`LIMIT` the number of open database connections",
		EnvVars: []string{"GOSH_DATABASE_CONNS"},
	},
	&cli.IntFlag{
		Name:    "database-conns-idle",
		Usage:   "`LIMIT` the number of idle database connections",
		EnvVars: []string{"GOSH_DATABASE_CONNS_IDLE"},
	},
	&cli.StringFlag{
		Name:    "download-token",
		Usage:   "optional static `SECRET` url token for file downloads",
		EnvVars: []string{"GOSH_DOWNLOAD_TOKEN"},
	},
	&cli.StringFlag{
		Name:    "preview-token",
		Usage:   "optional static `SECRET` url token for preview images and video streaming",
		EnvVars: []string{"GOSH_PREVIEW_TOKEN"},
	},
	&cli.BoolFlag{
		Name:    "prefork",
		Usage:   "spawn multiple Go processes listening on the same port",
		EnvVars: []string{"FIBER_PREFORK"},
	},
	&cli.StringFlag{
		Name:    "server-header",
		Usage:   "enable the 'Server: value' HTTP header",
		EnvVars: []string{"FIBER_SERVERHEADER"},
	},
	&cli.BoolFlag{
		Name:    "strict-routing",
		Usage:   "whether treats '/foo' and '/foo/' as different",
		EnvVars: []string{"FIBER_STRICTROUTING"},
	},
	&cli.BoolFlag{
		Name:    "case-sensitive",
		Usage:   "enables case sensitive routing",
		EnvVars: []string{"FIBER_CASESENSITIVE"},
	},
	&cli.BoolFlag{
		Name:    "immutable",
		Usage:   "relinquishes the 0-allocation promise in certain cases",
		EnvVars: []string{"FIBER_IMMUTABLE"},
	},
	&cli.BoolFlag{
		Name: "unescape-path",
		Usage: "converts all encoded characters in the route back " +
			"before setting the path for the context",
		EnvVars: []string{"FIBER_UNESCAPEPATH"},
	},
	&cli.BoolFlag{
		Name:    "etag",
		Usage:   "enable or disable ETag header generation",
		EnvVars: []string{"FIBER_ETAG"},
	},
	&cli.IntFlag{
		Name:    "body-limit",
		Usage:   "max body size that the server accepts",
		EnvVars: []string{"FIBER_BODYLIMIT"},
	},
	&cli.IntFlag{
		Name:    "concurrency",
		Usage:   "maximum number of concurrent connections",
		EnvVars: []string{"FIBER_CONCURRENCY"},
	},
	&cli.StringFlag{
		Name:    "views-layout",
		Usage:   "global layout for all template render",
		EnvVars: []string{"FIBER_VIEWS_LAYOUT"},
	},
	&cli.IntFlag{
		Name:    "read-timeout",
		Usage:   "amount of time allowed to read the full request including body",
		EnvVars: []string{"FIBER_READTIMEOUT"},
	},
	&cli.IntFlag{
		Name:    "idle-timeout",
		Usage:   "maximum amount of time to wait for the next request when keep-alive is enabled",
		EnvVars: []string{"FIBER_IDLETIMEOUT"},
	},
	&cli.IntFlag{
		Name:    "write-timeout",
		Usage:   "maximum duration before timing out writes of the response",
		EnvVars: []string{"FIBER_WRITETIMEOUT"},
	},
	&cli.IntFlag{
		Name:    "read-buffer-size",
		Usage:   "per-connection buffer size for requests' reading",
		EnvVars: []string{"FIBER_READBUFFERSIZE"},
	},
	&cli.IntFlag{
		Name:    "write-buffer-size",
		Usage:   "per-connection buffer size for responses' writing",
		EnvVars: []string{"FIBER_WRITEBUFFERSIZE"},
	},
	&cli.StringFlag{
		Name:    "compressed-file-suffix",
		Usage:   "adds suffix to the original file name",
		EnvVars: []string{"FIBER_COMPRESSEDFILESUFFIX"},
	},
	&cli.StringFlag{
		Name:    "proxy-header",
		Usage:   "enable c.IP() to return the value of the given header key",
		EnvVars: []string{"FIBER_PROXYHEADER"},
	},
	&cli.BoolFlag{
		Name:    "get-only",
		Usage:   "rejects all non-GET requests if set to true",
		EnvVars: []string{"FIBER_GETONLY"},
	},
	&cli.BoolFlag{
		Name:    "disable-keepalive",
		Usage:   "disables keep-alive connections",
		EnvVars: []string{"FIBER_DISABLEKEEPALIVE"},
	},
	&cli.BoolFlag{
		Name:    "disable-default-date",
		Usage:   "excludes the default date header from the response",
		EnvVars: []string{"FIBER_DISABLEDEFAULTDATE"},
	},
	&cli.BoolFlag{
		Name:    "disable-default-content-type",
		Usage:   "excludes the default Content-Type header from the response",
		EnvVars: []string{"FIBER_DISABLEDEFAULTCONTENTTYPE"},
	},
	&cli.BoolFlag{
		Name:    "disable-header-normalizing",
		Usage:   "disables header normalization",
		EnvVars: []string{"FIBER_DISABLEHEADERNORMALIZING"},
	},
	&cli.BoolFlag{
		Name:    "disable-startup-message",
		Usage:   "not print out the «Fiber» ASCII art and listening address",
		EnvVars: []string{"FIBER_DISABLESTARTUPMESSAGE"},
	},
	&cli.BoolFlag{
		Name:    "reduce-memory-usage",
		Usage:   "reduce memory usage",
		EnvVars: []string{"FIBER_REDUCEMEMORYUSAGE"},
	},
	&cli.BoolFlag{
		Name:    "enable-trusted-proxy-check",
		Usage:   "enable TrustedProxyCheck to prevent header spoofing",
		EnvVars: []string{"FIBER_REDUCEMEMORYUSAGE"},
	},
	&cli.StringFlag{
		Name:    "trusted-proxies",
		Usage:   "Trusted Proxies whitelist",
		EnvVars: []string{"FIBER_REDUCEMEMORYUSAGE"},
	},
}
