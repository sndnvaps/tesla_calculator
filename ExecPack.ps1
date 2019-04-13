# This PowerShell Script is write by sndnvaps@gmail.com
# what it can do Build the Application from code
# Pack && deploy it.
# Now just pack for Qt5.11.1
# if you want to pack for what you Qt version
# you must install goqt first
$QtVersion="qt5.11.1"
$AppName="tesla_calculator"
$AppVersion=$(git rev-parse --short HEAD)
$ReleaseDir="Release"
$DependDir="Depends"
$gopathTmp=$env:GOPATH

Function Build {
    new-item -type directory -path $ReleaseDir
    go build -ldflags "-H windowsgui" -o $ReleaseDir\$AppName.exe
}

Function Clean {
 if (Test-path $AppName-v$AppVersion-win-x86.zip) {
	   remove-item $AppName-v$AppVersion-win-x86.zip
	}

 if (Test-path $ReleaseDir) {
       remove-item $ReleaseDir -Recurse
	}
}

Function CopyDepend {
    copy-item conf -destination $ReleaseDir\ -Recurse
	copy-item $gopathTmp\src\github.com\visualfc\goqt\bin\qtdrv.ui.dll -destination $ReleaseDir
}

Function Deploy {
   windeployqt $ReleaseDir\qtdrv.ui.dll --dir $ReleaseDir
}

Function PackZip {
    Tools\7za.exe a $AppName-v$AppVersion-win-x86.zip $ReleaseDir\
}

Clean
Build
CopyDepend
Deploy
PackZip







