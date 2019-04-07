# This PowerShell Script is write by sndnvaps@gmail.com
# what it can do Build the Application from code
# Pack && deploy it.
# Now just pack for Qt5.11.1
# if you want to pack for what you Qt version
$QtVersion="qt5.11.1"
$AppName="tesla_calculator.exe"
$AppVersion="1.5.0"
$ReleaseDir="Release"
$DependDir="Depends"

Function Build {
    new-item -type directory -path $ReleaseDir
    go build -ldflags "-H windowsgui" -o $ReleaseDir\$AppName
}

Function Clean {
 if (Test-path $AppName-v$AppVersion-win-x86.zip) {
	remove-item $AppName-v$AppVersion-win-x86.zip
	}
 if (Test-path $AppName) {
     remove-item $AppName
	}
  if (Test-path $ReleaseDir) {
     remove-item $ReleaseDir -Recurse
	}
}

Function CopyDepend {
	new-item -type directory -path $ReleaseDir\platforms
	
	copy-item $DependDir\$QtVersion\plugins\win_x86\*.dll -destination $ReleaseDir\platforms
	copy-item $DependDir\$QtVersion\win_x86\*.dll -destination $ReleaseDir
}

Function PackZip {
    Tools\7za.exe a $AppName-v$AppVersion-win-x86.zip $ReleaseDir\
}

Clean
Build
CopyDepend
PackZip







