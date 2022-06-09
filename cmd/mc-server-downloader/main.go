package main
import (
    "log"
    "encoding/json"
    "fmt"
	"github.com/rmanders/minecraft-server-downloader/internal/utils"
    "github.com/rmanders/minecraft-server-downloader/internal/mojang"
)

func main() {

    var err error = nil

    // (1) Get the main manifest
    manifest_url := "https://launchermeta.mojang.com/mc/game/version_manifest.json"
    manifestJson, manifestErr := utils.GetJsonBytesFromUrl(manifest_url)
    if manifestErr != nil {
        log.Fatalln(manifestErr)
    }

    var minecraftVersions mojang.Versions
    err = json.Unmarshal(manifestJson, &minecraftVersions)
    if err != nil {
        log.Fatalln(err)
    }

    currentVersionNo := minecraftVersions.Latest.Release
    fmt.Println("current release: ", currentVersionNo)

    // (2) Get the current version info
    var currentVersionInfo mojang.VersionInfo
    for i := 0; i<len(minecraftVersions.Versions); i++ {
        if minecraftVersions.Versions[i].Id == currentVersionNo {
            currentVersionInfo = minecraftVersions.Versions[i]
            break
        }
    }
    // TODO: Handle version not found

    fmt.Println("current version Info:", currentVersionInfo)

    // (3) Get the package metadata for the current version
    packageJson, packageErr := utils.GetJsonBytesFromUrl(currentVersionInfo.Url)
    if packageErr != nil {
        log.Fatalln(packageErr)
    }

    var packageMetadata mojang.PackageMetadata
    err = json.Unmarshal(packageJson, &packageMetadata)
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println("Server Jar Url:", packageMetadata.Downloads.Server.Url)

    // (4) Download the file
    filename := fmt.Sprintf("minecraft-server-%s.jar", currentVersionNo)
    err = utils.DownloadFile(filename, packageMetadata.Downloads.Server.Url)
    if err != nil {
        log.Fatalln(err)
    }

    err = utils.CheckSha1(filename, packageMetadata.Downloads.Server.Sha1)
    if err != nil {
        log.Fatalln(err)
    }
}

