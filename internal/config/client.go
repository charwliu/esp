package config

import (
	"strings"

	"go.vixal.xyz/esp/pkg/fs"
	"go.vixal.xyz/esp/pkg/txt"
)

// ClientConfig represents HTTP client / Web UI config options.
type ClientConfig struct {
	Mode            string   `json:"mode,omitempty"`
	Name            string   `json:"name,omitempty"`
	Version         string   `json:"version,omitempty"`
	Copyright       string   `json:"copyright,omitempty"`
	Flags           string   `json:"flags,omitempty"`
	BaseUri         string   `json:"baseUri,omitempty"`
	StaticUri       string   `json:"staticUri,omitempty"`
	ApiUri          string   `json:"apiUri,omitempty"`
	ContentUri      string   `json:"contentUri,omitempty"`
	SiteUrl         string   `json:"siteUrl,omitempty"`
	SitePreview     string   `json:"sitePreview,omitempty"`
	SiteTitle       string   `json:"siteTitle,omitempty"`
	SiteCaption     string   `json:"siteCaption,omitempty"`
	SiteDescription string   `json:"siteDescription,omitempty"`
	SiteAuthor      string   `json:"siteAuthor,omitempty"`
	Debug           bool     `json:"debug,omitempty"`
	Test            bool     `json:"test,omitempty"`
	Demo            bool     `json:"demo,omitempty"`
	Sponsor         bool     `json:"sponsor,omitempty"`
	ReadOnly        bool     `json:"readonly,omitempty"`
	Public          bool     `json:"public,omitempty"`
	Experimental    bool     `json:"experimental,omitempty"`
	DownloadToken string              `json:"downloadToken,omitempty"`
	PreviewToken  string              `json:"previewToken,omitempty"`
	JSHash        string              `json:"jsHash,omitempty"`
	CSSHash       string              `json:"cssHash,omitempty"`
	ManifestHash  string              `json:"manifestHash,omitempty"`
	Settings      Settings            `json:"settings,omitempty"`
	Disable       ClientDisable       `json:"disable,omitempty"`
	Years         []int               `json:"years,omitempty"`
	Colors        []map[string]string `json:"colors,omitempty"`
	Categories    []CategoryLabel     `json:"categories,omitempty"`
	Clip          int                 `json:"clip,omitempty"`
	Server        RuntimeInfo         `json:"server,omitempty"`
}

// ClientDisable represents disabled client features a user can't turn back on.
type ClientDisable struct {
	Backups     bool `json:"backups"`
	WebDAV      bool `json:"webdav"`
	Settings    bool `json:"settings"`
	Places      bool `json:"places"`
	ExifTool    bool `json:"exiftool"`
	Darktable   bool `json:"darktable"`
	Rawtherapee bool `json:"rawtherapee"`
	Sips        bool `json:"sips"`
	HeifConvert bool `json:"heifconvert"`
	FFmpeg      bool `json:"ffmpeg"`
	TensorFlow  bool `json:"tensorflow"`
}


type CategoryLabel struct {
	LabelUID   string `json:"UID"`
	CustomSlug string `json:"Slug"`
	LabelName  string `json:"UserName"`
}


// Flags returns config flags as string slice.
func (c *Config) Flags() (flags []string) {
	if c.Public() {
		flags = append(flags, "public")
	}

	if c.Debug() {
		flags = append(flags, "debug")
	}

	if c.Experimental() {
		flags = append(flags, "experimental")
	}

	if c.ReadOnly() {
		flags = append(flags, "readonly")
	}

	if !c.DisableSettings() {
		flags = append(flags, "settings")
	}

	if !c.Settings().UI.Scrollbar {
		flags = append(flags, "hide-scrollbar")
	}

	return flags
}

// PublicConfig returns public client config options with as little information as possible.
func (c *Config) PublicConfig() ClientConfig {
	if c.Public() {
		return c.UserConfig()
	}

	settings := c.Settings()

	result := ClientConfig{
		Settings: Settings{
			UI:       settings.UI,
			Maps:     settings.Maps,
			Features: settings.Features,
			Share:    settings.Share,
		},
		Disable: ClientDisable{
			Backups:     true,
			WebDAV:      true,
			Settings:    c.DisableSettings(),
			Places:      c.DisablePlaces(),
			ExifTool:    true,
			TensorFlow:  true,
			Darktable:   true,
			Rawtherapee: true,
			Sips:        true,
			HeifConvert: true,
			FFmpeg:      true,
		},
		Flags:           strings.Join(c.Flags(), " "),
		Mode:            "public",
		Name:            c.Name(),
		BaseUri:         c.BaseUri(""),
		StaticUri:       c.StaticUri(),
		ApiUri:          c.ApiUri(),
		ContentUri:      c.ContentUri(),
		SiteUrl:         c.SiteUrl(),
		SitePreview:     c.SitePreview(),
		SiteTitle:       c.SiteTitle(),
		SiteCaption:     c.SiteCaption(),
		SiteDescription: c.SiteDescription(),
		SiteAuthor:      c.SiteAuthor(),
		Version:         c.Version(),
		Copyright:       c.Copyright(),
		Debug:           c.Debug(),
		Test:            c.Test(),
		Demo:            c.Demo(),
		Sponsor:         c.Sponsor(),
		ReadOnly:        c.ReadOnly(),
		Public:          c.Public(),
		Experimental:    c.Experimental(),
		JSHash:          fs.Checksum(c.BuildPath() + "/app.js"),
		CSSHash:         fs.Checksum(c.BuildPath() + "/app.css"),
		ManifestHash:    fs.Checksum(c.TemplatesPath() + "/manifest.json"),
		Clip:            txt.ClipDefault,
		PreviewToken:    "public",
		DownloadToken:   "public",
	}

	return result
}

// GuestConfig returns client config options for the sharing with guests.
func (c *Config) GuestConfig() ClientConfig {
	settings := c.Settings()

	result := ClientConfig{
		Settings: Settings{
			UI:       settings.UI,
			Maps:     settings.Maps,
			Features: settings.Features,
			Share:    settings.Share,
		},
		Disable: ClientDisable{
			Backups:     true,
			WebDAV:      c.DisableWebDAV(),
			Settings:    c.DisableSettings(),
			Places:      c.DisablePlaces(),
			ExifTool:    true,
			TensorFlow:  true,
			Darktable:   true,
			Rawtherapee: true,
			Sips:        true,
			HeifConvert: true,
			FFmpeg:      true,
		},
		Flags:           "readonly public shared",
		Mode:            "guest",
		Name:            c.Name(),
		BaseUri:         c.BaseUri(""),
		StaticUri:       c.StaticUri(),
		ApiUri:          c.ApiUri(),
		ContentUri:      c.ContentUri(),
		SiteUrl:         c.SiteUrl(),
		SitePreview:     c.SitePreview(),
		SiteTitle:       c.SiteTitle(),
		SiteCaption:     c.SiteCaption(),
		SiteDescription: c.SiteDescription(),
		SiteAuthor:      c.SiteAuthor(),
		Version:         c.Version(),
		Copyright:       c.Copyright(),
		Debug:           c.Debug(),
		Test:            c.Test(),
		Demo:            c.Demo(),
		Sponsor:         c.Sponsor(),
		ReadOnly:        true,
		Public:          true,
		Experimental:    false,
		DownloadToken:   c.DownloadToken(),
		PreviewToken:    c.PreviewToken(),
		JSHash:          fs.Checksum(c.BuildPath() + "/share.js"),
		CSSHash:         fs.Checksum(c.BuildPath() + "/share.css"),
		ManifestHash:    fs.Checksum(c.TemplatesPath() + "/manifest.json"),
		Clip:            txt.ClipDefault,
	}

	return result
}

// UserConfig returns client configuration options for registered users.
func (c *Config) UserConfig() ClientConfig {
	result := ClientConfig{
		Settings: *c.Settings(),
		Disable: ClientDisable{
			Backups:  c.DisableBackups(),
			WebDAV:   c.DisableWebDAV(),
			Settings: c.DisableSettings(),
			Places:   c.DisablePlaces(),
		},
		Flags:           strings.Join(c.Flags(), " "),
		Mode:            "user",
		Name:            c.Name(),
		BaseUri:         c.BaseUri(""),
		StaticUri:       c.StaticUri(),
		ApiUri:          c.ApiUri(),
		ContentUri:      c.ContentUri(),
		SiteUrl:         c.SiteUrl(),
		SitePreview:     c.SitePreview(),
		SiteTitle:       c.SiteTitle(),
		SiteCaption:     c.SiteCaption(),
		SiteDescription: c.SiteDescription(),
		SiteAuthor:      c.SiteAuthor(),
		Version:         c.Version(),
		Copyright:       c.Copyright(),
		Debug:           c.Debug(),
		Test:            c.Test(),
		Demo:            c.Demo(),
		Sponsor:         c.Sponsor(),
		ReadOnly:        c.ReadOnly(),
		Public:          c.Public(),
		Experimental:    c.Experimental(),
		DownloadToken:   c.DownloadToken(),
		PreviewToken:    c.PreviewToken(),
		JSHash:          fs.Checksum(c.BuildPath() + "/app.js"),
		CSSHash:         fs.Checksum(c.BuildPath() + "/app.css"),
		ManifestHash:    fs.Checksum(c.TemplatesPath() + "/manifest.json"),
		Clip:            txt.ClipDefault,
		Server:          NewRuntimeInfo(),
	}



	return result
}
