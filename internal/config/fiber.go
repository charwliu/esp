package config

import (
	"strconv"
	"time"
)

func (c *Config) TrustedProxies() []string {
	return c.options.TrustedProxies
}

func (c *Config) Prefork() bool {
	return c.options.Prefork
}

func (c *Config) ServerHeader() string {
	return c.options.ServerHeader
}

func (c *Config) StrictRouting() bool {
	return c.options.StrictRouting
}

func (c *Config) HasherDriver() string {
	return c.options.HasherDriver
}

func (c *Config) SessionLookup() string {
	return c.options.SessionLookup
}

func (c *Config) CaseSensitive() bool {
	return c.options.CaseSensitive
}

func (c *Config) GETOnly() bool {
	return c.options.GETOnly
}

func (c *Config) SessionHost() string {
	return c.options.SessionHost
}

func (c *Config) FiberViews() string {
	return c.options.FiberViews
}

func (c *Config) Immutable() bool {
	return c.options.Immutable
}

func (c *Config) SessionSecure() bool {
	return c.options.SessionSecure
}

func (c *Config) SessionDomain() string {
	return c.options.SessionDomain
}

func (c *Config) SessionSameSite() string {
	return c.options.SessionSameSite
}

func (c *Config) SessionExpiration() time.Duration {
	return c.options.SessionExpiration
}

func (c *Config) SessionGCInterval() time.Duration {
	return c.options.SessionGCInterval
}

func (c *Config) FiberViewsLayout() string {
	return c.options.FiberViewsLayout
}

func (c *Config) FiberViewsDelimsL() string {
	return c.options.FiberViewsDelimsL
}

func (c *Config) FiberViewsDelimsR() string {
	return c.options.FiberViewsDelimsR
}

func (c *Config) FiberViewsExtension() string {
	return c.options.FiberViewsExtension
}

func (c *Config) FiberViewsDebug() bool {
	return c.options.FiberViewsDebug
}

func (c *Config) FiberViewsReload() bool {
	return c.options.FiberViewsReload
}

func (c *Config) FiberViewsDirectory() string {
	return c.options.FiberViewsDirectory
}

func (c *Config) SetFiberViewsExtension(s string) {
	c.options.FiberViewsExtension = s
}

func (c *Config) UnescapePath() bool {
	return c.options.UnescapePath
}

func (c *Config) ETag() bool {
	return c.options.ETag
}

func (c *Config) BodyLimit() int {
	return c.options.BodyLimit
}

func (c *Config) Concurrency() int {
	return c.options.Concurrency
}

func (c *Config) ReadTimeout() time.Duration {
	return c.options.ReadTimeout
}

func (c *Config) WriteTimeout() time.Duration {
	return c.options.WriteTimeout
}

func (c *Config) IdleTimeout() time.Duration {
	return c.options.IdleTimeout
}

func (c *Config) ReadBufferSize() int {
	return c.options.ReadBufferSize
}

func (c *Config) WriteBufferSize() int {
	return c.options.WriteBufferSize
}

func (c *Config) CompressedFileSuffix() string {
	return c.options.CompressedFileSuffix
}

func (c *Config) ProxyHeader() string {
	return c.options.ProxyHeader
}

func (c *Config) DisableKeepalive() bool {
	return c.options.DisableKeepalive
}

func (c *Config) DisableDefaultDate() bool {
	return c.options.DisableDefaultDate
}

func (c *Config) DisableDefaultContentType() bool {
	return c.options.DisableDefaultContentType
}

func (c *Config) DisableHeaderNormalizing() bool {
	return c.options.DisableHeaderNormalizing
}

func (c *Config) DisableStartupMessage() bool {
	return c.options.DisableStartupMessage
}

func (c *Config) ReduceMemoryUsage() bool {
	return c.options.ReduceMemoryUsage
}

func (c *Config) SessionDatabase() string {
	return c.options.SessionDatabase
}

func (c *Config) SessionProvider() string {
	return c.options.SessionProvider
}

func (c *Config) HasherRounds() int {
	return c.options.HasherRounds
}

func (c *Config) SessionKeyPrefix() string {
	return c.options.SessionKeyPrefix
}

func (c *Config) SessionPort() int {
	return c.options.SessionPort
}

func (c *Config) HasherMemory() uint32 {
	return c.options.HasherMemory
}

func (c *Config) SessionUsername() string {
	return c.options.SessionUsername
}

func (c *Config) HasherSaltLength() uint32 {
	return c.options.HasherSaltLength
}

func (c *Config) HasherKeyLength() uint32 {
	return c.options.HasherKeyLength
}

func (c *Config) SessionTableName() string {
	return c.options.SessionTableName
}

func (c *Config) SessionPassword() string {
	return c.options.SessionPassword
}

func (c *Config) SessionDB() int {
	i, _ := strconv.Atoi(c.options.SessionDatabase)
	return i
}

func (c *Config) HasherIterations() uint32 {
	return c.options.HasherIterations
}

func (c *Config) HasherParallelism() uint8 {
	return c.options.HasherParallelism
}
