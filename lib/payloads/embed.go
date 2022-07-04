package payloads

import (
	"embed"
)

//go:embed behinder/java/*.class
var BehinderClassPayloads embed.FS

//go:embed behinder/php/*.php
var BehinderPhpPayloads embed.FS

//go:embed behinder/csharp/*.dll
var BehinderCsharpPayloads embed.FS

//go:embed behinder/asp/*.asp
var BehinderAspPayloads embed.FS

//go:embed godzilla/java/payloadv4.class
var GodzillaClassPayload []byte

//go:embed godzilla/java/plugins/*.class
var GodClassPluginsFiles embed.FS

//go:embed godzilla/php/payloadv4.php.txt
var GodzillaPhpPayload []byte

//go:embed godzilla/php/plugins/*.php
var GodPhpPluginsFiles embed.FS

//go:embed godzilla/csharp/payload.dll
var GodzillaCsharpPayload []byte

//go:embed godzilla/csharp/plugins/*.dll
var GodDllPluginsFiles embed.FS

//go:embed godzilla/asp/payload.asp.txt
var GodzillaAspPayload []byte

//go:embed godzilla/asp/plugins/*.asp
var GodAspPluginsFiles embed.FS
