package model

import "github.com/admpub/nging/v3/application/dbschema"

type CloudBackupExt struct {
	*dbschema.NgingCloudBackup
	Storage       *dbschema.NgingCloudStorage `db:"-,relation=id:dest_storage"`
	Watching      bool
	FullBackuping bool
}
