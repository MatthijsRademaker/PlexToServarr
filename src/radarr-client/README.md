# Go API client for swagger

Radarr API docs

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 3.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *{protocol}://{hostpath}*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AlternativeTitleApi* | [**ApiV3AlttitleGet**](docs/AlternativeTitleApi.md#apiv3alttitleget) | **Get** /api/v3/alttitle | 
*AlternativeTitleApi* | [**ApiV3AlttitleIdGet**](docs/AlternativeTitleApi.md#apiv3alttitleidget) | **Get** /api/v3/alttitle/{id} | 
*ApiInfoApi* | [**ApiGet**](docs/ApiInfoApi.md#apiget) | **Get** /api | 
*AuthenticationApi* | [**LoginPost**](docs/AuthenticationApi.md#loginpost) | **Post** /login | 
*AuthenticationApi* | [**LogoutGet**](docs/AuthenticationApi.md#logoutget) | **Get** /logout | 
*AutoTaggingApi* | [**ApiV3AutotaggingGet**](docs/AutoTaggingApi.md#apiv3autotaggingget) | **Get** /api/v3/autotagging | 
*AutoTaggingApi* | [**ApiV3AutotaggingIdDelete**](docs/AutoTaggingApi.md#apiv3autotaggingiddelete) | **Delete** /api/v3/autotagging/{id} | 
*AutoTaggingApi* | [**ApiV3AutotaggingIdGet**](docs/AutoTaggingApi.md#apiv3autotaggingidget) | **Get** /api/v3/autotagging/{id} | 
*AutoTaggingApi* | [**ApiV3AutotaggingIdPut**](docs/AutoTaggingApi.md#apiv3autotaggingidput) | **Put** /api/v3/autotagging/{id} | 
*AutoTaggingApi* | [**ApiV3AutotaggingPost**](docs/AutoTaggingApi.md#apiv3autotaggingpost) | **Post** /api/v3/autotagging | 
*AutoTaggingApi* | [**ApiV3AutotaggingSchemaGet**](docs/AutoTaggingApi.md#apiv3autotaggingschemaget) | **Get** /api/v3/autotagging/schema | 
*BackupApi* | [**ApiV3SystemBackupGet**](docs/BackupApi.md#apiv3systembackupget) | **Get** /api/v3/system/backup | 
*BackupApi* | [**ApiV3SystemBackupIdDelete**](docs/BackupApi.md#apiv3systembackupiddelete) | **Delete** /api/v3/system/backup/{id} | 
*BackupApi* | [**ApiV3SystemBackupRestoreIdPost**](docs/BackupApi.md#apiv3systembackuprestoreidpost) | **Post** /api/v3/system/backup/restore/{id} | 
*BackupApi* | [**ApiV3SystemBackupRestoreUploadPost**](docs/BackupApi.md#apiv3systembackuprestoreuploadpost) | **Post** /api/v3/system/backup/restore/upload | 
*BlocklistApi* | [**ApiV3BlocklistBulkDelete**](docs/BlocklistApi.md#apiv3blocklistbulkdelete) | **Delete** /api/v3/blocklist/bulk | 
*BlocklistApi* | [**ApiV3BlocklistGet**](docs/BlocklistApi.md#apiv3blocklistget) | **Get** /api/v3/blocklist | 
*BlocklistApi* | [**ApiV3BlocklistIdDelete**](docs/BlocklistApi.md#apiv3blocklistiddelete) | **Delete** /api/v3/blocklist/{id} | 
*BlocklistApi* | [**ApiV3BlocklistMovieGet**](docs/BlocklistApi.md#apiv3blocklistmovieget) | **Get** /api/v3/blocklist/movie | 
*CalendarApi* | [**ApiV3CalendarGet**](docs/CalendarApi.md#apiv3calendarget) | **Get** /api/v3/calendar | 
*CalendarFeedApi* | [**FeedV3CalendarRadarrIcsGet**](docs/CalendarFeedApi.md#feedv3calendarradarricsget) | **Get** /feed/v3/calendar/radarr.ics | 
*CollectionApi* | [**ApiV3CollectionGet**](docs/CollectionApi.md#apiv3collectionget) | **Get** /api/v3/collection | 
*CollectionApi* | [**ApiV3CollectionIdGet**](docs/CollectionApi.md#apiv3collectionidget) | **Get** /api/v3/collection/{id} | 
*CollectionApi* | [**ApiV3CollectionIdPut**](docs/CollectionApi.md#apiv3collectionidput) | **Put** /api/v3/collection/{id} | 
*CollectionApi* | [**ApiV3CollectionPut**](docs/CollectionApi.md#apiv3collectionput) | **Put** /api/v3/collection | 
*CommandApi* | [**ApiV3CommandGet**](docs/CommandApi.md#apiv3commandget) | **Get** /api/v3/command | 
*CommandApi* | [**ApiV3CommandIdDelete**](docs/CommandApi.md#apiv3commandiddelete) | **Delete** /api/v3/command/{id} | 
*CommandApi* | [**ApiV3CommandIdGet**](docs/CommandApi.md#apiv3commandidget) | **Get** /api/v3/command/{id} | 
*CommandApi* | [**ApiV3CommandPost**](docs/CommandApi.md#apiv3commandpost) | **Post** /api/v3/command | 
*CreditApi* | [**ApiV3CreditGet**](docs/CreditApi.md#apiv3creditget) | **Get** /api/v3/credit | 
*CreditApi* | [**ApiV3CreditIdGet**](docs/CreditApi.md#apiv3creditidget) | **Get** /api/v3/credit/{id} | 
*CustomFilterApi* | [**ApiV3CustomfilterGet**](docs/CustomFilterApi.md#apiv3customfilterget) | **Get** /api/v3/customfilter | 
*CustomFilterApi* | [**ApiV3CustomfilterIdDelete**](docs/CustomFilterApi.md#apiv3customfilteriddelete) | **Delete** /api/v3/customfilter/{id} | 
*CustomFilterApi* | [**ApiV3CustomfilterIdGet**](docs/CustomFilterApi.md#apiv3customfilteridget) | **Get** /api/v3/customfilter/{id} | 
*CustomFilterApi* | [**ApiV3CustomfilterIdPut**](docs/CustomFilterApi.md#apiv3customfilteridput) | **Put** /api/v3/customfilter/{id} | 
*CustomFilterApi* | [**ApiV3CustomfilterPost**](docs/CustomFilterApi.md#apiv3customfilterpost) | **Post** /api/v3/customfilter | 
*CustomFormatApi* | [**ApiV3CustomformatBulkDelete**](docs/CustomFormatApi.md#apiv3customformatbulkdelete) | **Delete** /api/v3/customformat/bulk | 
*CustomFormatApi* | [**ApiV3CustomformatBulkPut**](docs/CustomFormatApi.md#apiv3customformatbulkput) | **Put** /api/v3/customformat/bulk | 
*CustomFormatApi* | [**ApiV3CustomformatGet**](docs/CustomFormatApi.md#apiv3customformatget) | **Get** /api/v3/customformat | 
*CustomFormatApi* | [**ApiV3CustomformatIdDelete**](docs/CustomFormatApi.md#apiv3customformatiddelete) | **Delete** /api/v3/customformat/{id} | 
*CustomFormatApi* | [**ApiV3CustomformatIdGet**](docs/CustomFormatApi.md#apiv3customformatidget) | **Get** /api/v3/customformat/{id} | 
*CustomFormatApi* | [**ApiV3CustomformatIdPut**](docs/CustomFormatApi.md#apiv3customformatidput) | **Put** /api/v3/customformat/{id} | 
*CustomFormatApi* | [**ApiV3CustomformatPost**](docs/CustomFormatApi.md#apiv3customformatpost) | **Post** /api/v3/customformat | 
*CustomFormatApi* | [**ApiV3CustomformatSchemaGet**](docs/CustomFormatApi.md#apiv3customformatschemaget) | **Get** /api/v3/customformat/schema | 
*CutoffApi* | [**ApiV3WantedCutoffGet**](docs/CutoffApi.md#apiv3wantedcutoffget) | **Get** /api/v3/wanted/cutoff | 
*DelayProfileApi* | [**ApiV3DelayprofileGet**](docs/DelayProfileApi.md#apiv3delayprofileget) | **Get** /api/v3/delayprofile | 
*DelayProfileApi* | [**ApiV3DelayprofileIdDelete**](docs/DelayProfileApi.md#apiv3delayprofileiddelete) | **Delete** /api/v3/delayprofile/{id} | 
*DelayProfileApi* | [**ApiV3DelayprofileIdGet**](docs/DelayProfileApi.md#apiv3delayprofileidget) | **Get** /api/v3/delayprofile/{id} | 
*DelayProfileApi* | [**ApiV3DelayprofileIdPut**](docs/DelayProfileApi.md#apiv3delayprofileidput) | **Put** /api/v3/delayprofile/{id} | 
*DelayProfileApi* | [**ApiV3DelayprofilePost**](docs/DelayProfileApi.md#apiv3delayprofilepost) | **Post** /api/v3/delayprofile | 
*DelayProfileApi* | [**ApiV3DelayprofileReorderIdPut**](docs/DelayProfileApi.md#apiv3delayprofilereorderidput) | **Put** /api/v3/delayprofile/reorder/{id} | 
*DiskSpaceApi* | [**ApiV3DiskspaceGet**](docs/DiskSpaceApi.md#apiv3diskspaceget) | **Get** /api/v3/diskspace | 
*DownloadClientApi* | [**ApiV3DownloadclientActionNamePost**](docs/DownloadClientApi.md#apiv3downloadclientactionnamepost) | **Post** /api/v3/downloadclient/action/{name} | 
*DownloadClientApi* | [**ApiV3DownloadclientBulkDelete**](docs/DownloadClientApi.md#apiv3downloadclientbulkdelete) | **Delete** /api/v3/downloadclient/bulk | 
*DownloadClientApi* | [**ApiV3DownloadclientBulkPut**](docs/DownloadClientApi.md#apiv3downloadclientbulkput) | **Put** /api/v3/downloadclient/bulk | 
*DownloadClientApi* | [**ApiV3DownloadclientGet**](docs/DownloadClientApi.md#apiv3downloadclientget) | **Get** /api/v3/downloadclient | 
*DownloadClientApi* | [**ApiV3DownloadclientIdDelete**](docs/DownloadClientApi.md#apiv3downloadclientiddelete) | **Delete** /api/v3/downloadclient/{id} | 
*DownloadClientApi* | [**ApiV3DownloadclientIdGet**](docs/DownloadClientApi.md#apiv3downloadclientidget) | **Get** /api/v3/downloadclient/{id} | 
*DownloadClientApi* | [**ApiV3DownloadclientIdPut**](docs/DownloadClientApi.md#apiv3downloadclientidput) | **Put** /api/v3/downloadclient/{id} | 
*DownloadClientApi* | [**ApiV3DownloadclientPost**](docs/DownloadClientApi.md#apiv3downloadclientpost) | **Post** /api/v3/downloadclient | 
*DownloadClientApi* | [**ApiV3DownloadclientSchemaGet**](docs/DownloadClientApi.md#apiv3downloadclientschemaget) | **Get** /api/v3/downloadclient/schema | 
*DownloadClientApi* | [**ApiV3DownloadclientTestPost**](docs/DownloadClientApi.md#apiv3downloadclienttestpost) | **Post** /api/v3/downloadclient/test | 
*DownloadClientApi* | [**ApiV3DownloadclientTestallPost**](docs/DownloadClientApi.md#apiv3downloadclienttestallpost) | **Post** /api/v3/downloadclient/testall | 
*DownloadClientConfigApi* | [**ApiV3ConfigDownloadclientGet**](docs/DownloadClientConfigApi.md#apiv3configdownloadclientget) | **Get** /api/v3/config/downloadclient | 
*DownloadClientConfigApi* | [**ApiV3ConfigDownloadclientIdGet**](docs/DownloadClientConfigApi.md#apiv3configdownloadclientidget) | **Get** /api/v3/config/downloadclient/{id} | 
*DownloadClientConfigApi* | [**ApiV3ConfigDownloadclientIdPut**](docs/DownloadClientConfigApi.md#apiv3configdownloadclientidput) | **Put** /api/v3/config/downloadclient/{id} | 
*ExtraFileApi* | [**ApiV3ExtrafileGet**](docs/ExtraFileApi.md#apiv3extrafileget) | **Get** /api/v3/extrafile | 
*FileSystemApi* | [**ApiV3FilesystemGet**](docs/FileSystemApi.md#apiv3filesystemget) | **Get** /api/v3/filesystem | 
*FileSystemApi* | [**ApiV3FilesystemMediafilesGet**](docs/FileSystemApi.md#apiv3filesystemmediafilesget) | **Get** /api/v3/filesystem/mediafiles | 
*FileSystemApi* | [**ApiV3FilesystemTypeGet**](docs/FileSystemApi.md#apiv3filesystemtypeget) | **Get** /api/v3/filesystem/type | 
*HealthApi* | [**ApiV3HealthGet**](docs/HealthApi.md#apiv3healthget) | **Get** /api/v3/health | 
*HistoryApi* | [**ApiV3HistoryFailedIdPost**](docs/HistoryApi.md#apiv3historyfailedidpost) | **Post** /api/v3/history/failed/{id} | 
*HistoryApi* | [**ApiV3HistoryGet**](docs/HistoryApi.md#apiv3historyget) | **Get** /api/v3/history | 
*HistoryApi* | [**ApiV3HistoryMovieGet**](docs/HistoryApi.md#apiv3historymovieget) | **Get** /api/v3/history/movie | 
*HistoryApi* | [**ApiV3HistorySinceGet**](docs/HistoryApi.md#apiv3historysinceget) | **Get** /api/v3/history/since | 
*HostConfigApi* | [**ApiV3ConfigHostGet**](docs/HostConfigApi.md#apiv3confighostget) | **Get** /api/v3/config/host | 
*HostConfigApi* | [**ApiV3ConfigHostIdGet**](docs/HostConfigApi.md#apiv3confighostidget) | **Get** /api/v3/config/host/{id} | 
*HostConfigApi* | [**ApiV3ConfigHostIdPut**](docs/HostConfigApi.md#apiv3confighostidput) | **Put** /api/v3/config/host/{id} | 
*ImportListApi* | [**ApiV3ImportlistActionNamePost**](docs/ImportListApi.md#apiv3importlistactionnamepost) | **Post** /api/v3/importlist/action/{name} | 
*ImportListApi* | [**ApiV3ImportlistBulkDelete**](docs/ImportListApi.md#apiv3importlistbulkdelete) | **Delete** /api/v3/importlist/bulk | 
*ImportListApi* | [**ApiV3ImportlistBulkPut**](docs/ImportListApi.md#apiv3importlistbulkput) | **Put** /api/v3/importlist/bulk | 
*ImportListApi* | [**ApiV3ImportlistGet**](docs/ImportListApi.md#apiv3importlistget) | **Get** /api/v3/importlist | 
*ImportListApi* | [**ApiV3ImportlistIdDelete**](docs/ImportListApi.md#apiv3importlistiddelete) | **Delete** /api/v3/importlist/{id} | 
*ImportListApi* | [**ApiV3ImportlistIdGet**](docs/ImportListApi.md#apiv3importlistidget) | **Get** /api/v3/importlist/{id} | 
*ImportListApi* | [**ApiV3ImportlistIdPut**](docs/ImportListApi.md#apiv3importlistidput) | **Put** /api/v3/importlist/{id} | 
*ImportListApi* | [**ApiV3ImportlistPost**](docs/ImportListApi.md#apiv3importlistpost) | **Post** /api/v3/importlist | 
*ImportListApi* | [**ApiV3ImportlistSchemaGet**](docs/ImportListApi.md#apiv3importlistschemaget) | **Get** /api/v3/importlist/schema | 
*ImportListApi* | [**ApiV3ImportlistTestPost**](docs/ImportListApi.md#apiv3importlisttestpost) | **Post** /api/v3/importlist/test | 
*ImportListApi* | [**ApiV3ImportlistTestallPost**](docs/ImportListApi.md#apiv3importlisttestallpost) | **Post** /api/v3/importlist/testall | 
*ImportListConfigApi* | [**ApiV3ConfigImportlistGet**](docs/ImportListConfigApi.md#apiv3configimportlistget) | **Get** /api/v3/config/importlist | 
*ImportListConfigApi* | [**ApiV3ConfigImportlistIdGet**](docs/ImportListConfigApi.md#apiv3configimportlistidget) | **Get** /api/v3/config/importlist/{id} | 
*ImportListConfigApi* | [**ApiV3ConfigImportlistIdPut**](docs/ImportListConfigApi.md#apiv3configimportlistidput) | **Put** /api/v3/config/importlist/{id} | 
*ImportListExclusionApi* | [**ApiV3ExclusionsBulkDelete**](docs/ImportListExclusionApi.md#apiv3exclusionsbulkdelete) | **Delete** /api/v3/exclusions/bulk | 
*ImportListExclusionApi* | [**ApiV3ExclusionsBulkPost**](docs/ImportListExclusionApi.md#apiv3exclusionsbulkpost) | **Post** /api/v3/exclusions/bulk | 
*ImportListExclusionApi* | [**ApiV3ExclusionsGet**](docs/ImportListExclusionApi.md#apiv3exclusionsget) | **Get** /api/v3/exclusions | 
*ImportListExclusionApi* | [**ApiV3ExclusionsIdDelete**](docs/ImportListExclusionApi.md#apiv3exclusionsiddelete) | **Delete** /api/v3/exclusions/{id} | 
*ImportListExclusionApi* | [**ApiV3ExclusionsIdGet**](docs/ImportListExclusionApi.md#apiv3exclusionsidget) | **Get** /api/v3/exclusions/{id} | 
*ImportListExclusionApi* | [**ApiV3ExclusionsIdPut**](docs/ImportListExclusionApi.md#apiv3exclusionsidput) | **Put** /api/v3/exclusions/{id} | 
*ImportListExclusionApi* | [**ApiV3ExclusionsPagedGet**](docs/ImportListExclusionApi.md#apiv3exclusionspagedget) | **Get** /api/v3/exclusions/paged | 
*ImportListExclusionApi* | [**ApiV3ExclusionsPost**](docs/ImportListExclusionApi.md#apiv3exclusionspost) | **Post** /api/v3/exclusions | 
*ImportListMoviesApi* | [**ApiV3ImportlistMovieGet**](docs/ImportListMoviesApi.md#apiv3importlistmovieget) | **Get** /api/v3/importlist/movie | 
*ImportListMoviesApi* | [**ApiV3ImportlistMoviePost**](docs/ImportListMoviesApi.md#apiv3importlistmoviepost) | **Post** /api/v3/importlist/movie | 
*IndexerApi* | [**ApiV3IndexerActionNamePost**](docs/IndexerApi.md#apiv3indexeractionnamepost) | **Post** /api/v3/indexer/action/{name} | 
*IndexerApi* | [**ApiV3IndexerBulkDelete**](docs/IndexerApi.md#apiv3indexerbulkdelete) | **Delete** /api/v3/indexer/bulk | 
*IndexerApi* | [**ApiV3IndexerBulkPut**](docs/IndexerApi.md#apiv3indexerbulkput) | **Put** /api/v3/indexer/bulk | 
*IndexerApi* | [**ApiV3IndexerGet**](docs/IndexerApi.md#apiv3indexerget) | **Get** /api/v3/indexer | 
*IndexerApi* | [**ApiV3IndexerIdDelete**](docs/IndexerApi.md#apiv3indexeriddelete) | **Delete** /api/v3/indexer/{id} | 
*IndexerApi* | [**ApiV3IndexerIdGet**](docs/IndexerApi.md#apiv3indexeridget) | **Get** /api/v3/indexer/{id} | 
*IndexerApi* | [**ApiV3IndexerIdPut**](docs/IndexerApi.md#apiv3indexeridput) | **Put** /api/v3/indexer/{id} | 
*IndexerApi* | [**ApiV3IndexerPost**](docs/IndexerApi.md#apiv3indexerpost) | **Post** /api/v3/indexer | 
*IndexerApi* | [**ApiV3IndexerSchemaGet**](docs/IndexerApi.md#apiv3indexerschemaget) | **Get** /api/v3/indexer/schema | 
*IndexerApi* | [**ApiV3IndexerTestPost**](docs/IndexerApi.md#apiv3indexertestpost) | **Post** /api/v3/indexer/test | 
*IndexerApi* | [**ApiV3IndexerTestallPost**](docs/IndexerApi.md#apiv3indexertestallpost) | **Post** /api/v3/indexer/testall | 
*IndexerConfigApi* | [**ApiV3ConfigIndexerGet**](docs/IndexerConfigApi.md#apiv3configindexerget) | **Get** /api/v3/config/indexer | 
*IndexerConfigApi* | [**ApiV3ConfigIndexerIdGet**](docs/IndexerConfigApi.md#apiv3configindexeridget) | **Get** /api/v3/config/indexer/{id} | 
*IndexerConfigApi* | [**ApiV3ConfigIndexerIdPut**](docs/IndexerConfigApi.md#apiv3configindexeridput) | **Put** /api/v3/config/indexer/{id} | 
*IndexerFlagApi* | [**ApiV3IndexerflagGet**](docs/IndexerFlagApi.md#apiv3indexerflagget) | **Get** /api/v3/indexerflag | 
*LanguageApi* | [**ApiV3LanguageGet**](docs/LanguageApi.md#apiv3languageget) | **Get** /api/v3/language | 
*LanguageApi* | [**ApiV3LanguageIdGet**](docs/LanguageApi.md#apiv3languageidget) | **Get** /api/v3/language/{id} | 
*LocalizationApi* | [**ApiV3LocalizationGet**](docs/LocalizationApi.md#apiv3localizationget) | **Get** /api/v3/localization | 
*LocalizationApi* | [**ApiV3LocalizationLanguageGet**](docs/LocalizationApi.md#apiv3localizationlanguageget) | **Get** /api/v3/localization/language | 
*LogApi* | [**ApiV3LogGet**](docs/LogApi.md#apiv3logget) | **Get** /api/v3/log | 
*LogFileApi* | [**ApiV3LogFileFilenameGet**](docs/LogFileApi.md#apiv3logfilefilenameget) | **Get** /api/v3/log/file/{filename} | 
*LogFileApi* | [**ApiV3LogFileGet**](docs/LogFileApi.md#apiv3logfileget) | **Get** /api/v3/log/file | 
*ManualImportApi* | [**ApiV3ManualimportGet**](docs/ManualImportApi.md#apiv3manualimportget) | **Get** /api/v3/manualimport | 
*ManualImportApi* | [**ApiV3ManualimportPost**](docs/ManualImportApi.md#apiv3manualimportpost) | **Post** /api/v3/manualimport | 
*MediaCoverApi* | [**ApiV3MediacoverMovieIdFilenameGet**](docs/MediaCoverApi.md#apiv3mediacovermovieidfilenameget) | **Get** /api/v3/mediacover/{movieId}/{filename} | 
*MediaManagementConfigApi* | [**ApiV3ConfigMediamanagementGet**](docs/MediaManagementConfigApi.md#apiv3configmediamanagementget) | **Get** /api/v3/config/mediamanagement | 
*MediaManagementConfigApi* | [**ApiV3ConfigMediamanagementIdGet**](docs/MediaManagementConfigApi.md#apiv3configmediamanagementidget) | **Get** /api/v3/config/mediamanagement/{id} | 
*MediaManagementConfigApi* | [**ApiV3ConfigMediamanagementIdPut**](docs/MediaManagementConfigApi.md#apiv3configmediamanagementidput) | **Put** /api/v3/config/mediamanagement/{id} | 
*MetadataApi* | [**ApiV3MetadataActionNamePost**](docs/MetadataApi.md#apiv3metadataactionnamepost) | **Post** /api/v3/metadata/action/{name} | 
*MetadataApi* | [**ApiV3MetadataGet**](docs/MetadataApi.md#apiv3metadataget) | **Get** /api/v3/metadata | 
*MetadataApi* | [**ApiV3MetadataIdDelete**](docs/MetadataApi.md#apiv3metadataiddelete) | **Delete** /api/v3/metadata/{id} | 
*MetadataApi* | [**ApiV3MetadataIdGet**](docs/MetadataApi.md#apiv3metadataidget) | **Get** /api/v3/metadata/{id} | 
*MetadataApi* | [**ApiV3MetadataIdPut**](docs/MetadataApi.md#apiv3metadataidput) | **Put** /api/v3/metadata/{id} | 
*MetadataApi* | [**ApiV3MetadataPost**](docs/MetadataApi.md#apiv3metadatapost) | **Post** /api/v3/metadata | 
*MetadataApi* | [**ApiV3MetadataSchemaGet**](docs/MetadataApi.md#apiv3metadataschemaget) | **Get** /api/v3/metadata/schema | 
*MetadataApi* | [**ApiV3MetadataTestPost**](docs/MetadataApi.md#apiv3metadatatestpost) | **Post** /api/v3/metadata/test | 
*MetadataApi* | [**ApiV3MetadataTestallPost**](docs/MetadataApi.md#apiv3metadatatestallpost) | **Post** /api/v3/metadata/testall | 
*MetadataConfigApi* | [**ApiV3ConfigMetadataGet**](docs/MetadataConfigApi.md#apiv3configmetadataget) | **Get** /api/v3/config/metadata | 
*MetadataConfigApi* | [**ApiV3ConfigMetadataIdGet**](docs/MetadataConfigApi.md#apiv3configmetadataidget) | **Get** /api/v3/config/metadata/{id} | 
*MetadataConfigApi* | [**ApiV3ConfigMetadataIdPut**](docs/MetadataConfigApi.md#apiv3configmetadataidput) | **Put** /api/v3/config/metadata/{id} | 
*MissingApi* | [**ApiV3WantedMissingGet**](docs/MissingApi.md#apiv3wantedmissingget) | **Get** /api/v3/wanted/missing | 
*MovieApi* | [**ApiV3MovieGet**](docs/MovieApi.md#apiv3movieget) | **Get** /api/v3/movie | 
*MovieApi* | [**ApiV3MovieIdDelete**](docs/MovieApi.md#apiv3movieiddelete) | **Delete** /api/v3/movie/{id} | 
*MovieApi* | [**ApiV3MovieIdGet**](docs/MovieApi.md#apiv3movieidget) | **Get** /api/v3/movie/{id} | 
*MovieApi* | [**ApiV3MovieIdPut**](docs/MovieApi.md#apiv3movieidput) | **Put** /api/v3/movie/{id} | 
*MovieApi* | [**ApiV3MoviePost**](docs/MovieApi.md#apiv3moviepost) | **Post** /api/v3/movie | 
*MovieEditorApi* | [**ApiV3MovieEditorDelete**](docs/MovieEditorApi.md#apiv3movieeditordelete) | **Delete** /api/v3/movie/editor | 
*MovieEditorApi* | [**ApiV3MovieEditorPut**](docs/MovieEditorApi.md#apiv3movieeditorput) | **Put** /api/v3/movie/editor | 
*MovieFileApi* | [**ApiV3MoviefileBulkDelete**](docs/MovieFileApi.md#apiv3moviefilebulkdelete) | **Delete** /api/v3/moviefile/bulk | 
*MovieFileApi* | [**ApiV3MoviefileEditorPut**](docs/MovieFileApi.md#apiv3moviefileeditorput) | **Put** /api/v3/moviefile/editor | 
*MovieFileApi* | [**ApiV3MoviefileGet**](docs/MovieFileApi.md#apiv3moviefileget) | **Get** /api/v3/moviefile | 
*MovieFileApi* | [**ApiV3MoviefileIdDelete**](docs/MovieFileApi.md#apiv3moviefileiddelete) | **Delete** /api/v3/moviefile/{id} | 
*MovieFileApi* | [**ApiV3MoviefileIdGet**](docs/MovieFileApi.md#apiv3moviefileidget) | **Get** /api/v3/moviefile/{id} | 
*MovieFileApi* | [**ApiV3MoviefileIdPut**](docs/MovieFileApi.md#apiv3moviefileidput) | **Put** /api/v3/moviefile/{id} | 
*MovieImportApi* | [**ApiV3MovieImportPost**](docs/MovieImportApi.md#apiv3movieimportpost) | **Post** /api/v3/movie/import | 
*MovieLookupApi* | [**ApiV3MovieLookupGet**](docs/MovieLookupApi.md#apiv3movielookupget) | **Get** /api/v3/movie/lookup | 
*MovieLookupApi* | [**ApiV3MovieLookupImdbGet**](docs/MovieLookupApi.md#apiv3movielookupimdbget) | **Get** /api/v3/movie/lookup/imdb | 
*MovieLookupApi* | [**ApiV3MovieLookupTmdbGet**](docs/MovieLookupApi.md#apiv3movielookuptmdbget) | **Get** /api/v3/movie/lookup/tmdb | 
*NamingConfigApi* | [**ApiV3ConfigNamingExamplesGet**](docs/NamingConfigApi.md#apiv3confignamingexamplesget) | **Get** /api/v3/config/naming/examples | 
*NamingConfigApi* | [**ApiV3ConfigNamingGet**](docs/NamingConfigApi.md#apiv3confignamingget) | **Get** /api/v3/config/naming | 
*NamingConfigApi* | [**ApiV3ConfigNamingIdGet**](docs/NamingConfigApi.md#apiv3confignamingidget) | **Get** /api/v3/config/naming/{id} | 
*NamingConfigApi* | [**ApiV3ConfigNamingIdPut**](docs/NamingConfigApi.md#apiv3confignamingidput) | **Put** /api/v3/config/naming/{id} | 
*NotificationApi* | [**ApiV3NotificationActionNamePost**](docs/NotificationApi.md#apiv3notificationactionnamepost) | **Post** /api/v3/notification/action/{name} | 
*NotificationApi* | [**ApiV3NotificationGet**](docs/NotificationApi.md#apiv3notificationget) | **Get** /api/v3/notification | 
*NotificationApi* | [**ApiV3NotificationIdDelete**](docs/NotificationApi.md#apiv3notificationiddelete) | **Delete** /api/v3/notification/{id} | 
*NotificationApi* | [**ApiV3NotificationIdGet**](docs/NotificationApi.md#apiv3notificationidget) | **Get** /api/v3/notification/{id} | 
*NotificationApi* | [**ApiV3NotificationIdPut**](docs/NotificationApi.md#apiv3notificationidput) | **Put** /api/v3/notification/{id} | 
*NotificationApi* | [**ApiV3NotificationPost**](docs/NotificationApi.md#apiv3notificationpost) | **Post** /api/v3/notification | 
*NotificationApi* | [**ApiV3NotificationSchemaGet**](docs/NotificationApi.md#apiv3notificationschemaget) | **Get** /api/v3/notification/schema | 
*NotificationApi* | [**ApiV3NotificationTestPost**](docs/NotificationApi.md#apiv3notificationtestpost) | **Post** /api/v3/notification/test | 
*NotificationApi* | [**ApiV3NotificationTestallPost**](docs/NotificationApi.md#apiv3notificationtestallpost) | **Post** /api/v3/notification/testall | 
*ParseApi* | [**ApiV3ParseGet**](docs/ParseApi.md#apiv3parseget) | **Get** /api/v3/parse | 
*PingApi* | [**PingGet**](docs/PingApi.md#pingget) | **Get** /ping | 
*PingApi* | [**PingHead**](docs/PingApi.md#pinghead) | **Head** /ping | 
*QualityDefinitionApi* | [**ApiV3QualitydefinitionGet**](docs/QualityDefinitionApi.md#apiv3qualitydefinitionget) | **Get** /api/v3/qualitydefinition | 
*QualityDefinitionApi* | [**ApiV3QualitydefinitionIdGet**](docs/QualityDefinitionApi.md#apiv3qualitydefinitionidget) | **Get** /api/v3/qualitydefinition/{id} | 
*QualityDefinitionApi* | [**ApiV3QualitydefinitionIdPut**](docs/QualityDefinitionApi.md#apiv3qualitydefinitionidput) | **Put** /api/v3/qualitydefinition/{id} | 
*QualityDefinitionApi* | [**ApiV3QualitydefinitionUpdatePut**](docs/QualityDefinitionApi.md#apiv3qualitydefinitionupdateput) | **Put** /api/v3/qualitydefinition/update | 
*QualityProfileApi* | [**ApiV3QualityprofileGet**](docs/QualityProfileApi.md#apiv3qualityprofileget) | **Get** /api/v3/qualityprofile | 
*QualityProfileApi* | [**ApiV3QualityprofileIdDelete**](docs/QualityProfileApi.md#apiv3qualityprofileiddelete) | **Delete** /api/v3/qualityprofile/{id} | 
*QualityProfileApi* | [**ApiV3QualityprofileIdGet**](docs/QualityProfileApi.md#apiv3qualityprofileidget) | **Get** /api/v3/qualityprofile/{id} | 
*QualityProfileApi* | [**ApiV3QualityprofileIdPut**](docs/QualityProfileApi.md#apiv3qualityprofileidput) | **Put** /api/v3/qualityprofile/{id} | 
*QualityProfileApi* | [**ApiV3QualityprofilePost**](docs/QualityProfileApi.md#apiv3qualityprofilepost) | **Post** /api/v3/qualityprofile | 
*QualityProfileSchemaApi* | [**ApiV3QualityprofileSchemaGet**](docs/QualityProfileSchemaApi.md#apiv3qualityprofileschemaget) | **Get** /api/v3/qualityprofile/schema | 
*QueueApi* | [**ApiV3QueueBulkDelete**](docs/QueueApi.md#apiv3queuebulkdelete) | **Delete** /api/v3/queue/bulk | 
*QueueApi* | [**ApiV3QueueGet**](docs/QueueApi.md#apiv3queueget) | **Get** /api/v3/queue | 
*QueueApi* | [**ApiV3QueueIdDelete**](docs/QueueApi.md#apiv3queueiddelete) | **Delete** /api/v3/queue/{id} | 
*QueueActionApi* | [**ApiV3QueueGrabBulkPost**](docs/QueueActionApi.md#apiv3queuegrabbulkpost) | **Post** /api/v3/queue/grab/bulk | 
*QueueActionApi* | [**ApiV3QueueGrabIdPost**](docs/QueueActionApi.md#apiv3queuegrabidpost) | **Post** /api/v3/queue/grab/{id} | 
*QueueDetailsApi* | [**ApiV3QueueDetailsGet**](docs/QueueDetailsApi.md#apiv3queuedetailsget) | **Get** /api/v3/queue/details | 
*QueueStatusApi* | [**ApiV3QueueStatusGet**](docs/QueueStatusApi.md#apiv3queuestatusget) | **Get** /api/v3/queue/status | 
*ReleaseApi* | [**ApiV3ReleaseGet**](docs/ReleaseApi.md#apiv3releaseget) | **Get** /api/v3/release | 
*ReleaseApi* | [**ApiV3ReleasePost**](docs/ReleaseApi.md#apiv3releasepost) | **Post** /api/v3/release | 
*ReleaseProfileApi* | [**ApiV3ReleaseprofileGet**](docs/ReleaseProfileApi.md#apiv3releaseprofileget) | **Get** /api/v3/releaseprofile | 
*ReleaseProfileApi* | [**ApiV3ReleaseprofileIdDelete**](docs/ReleaseProfileApi.md#apiv3releaseprofileiddelete) | **Delete** /api/v3/releaseprofile/{id} | 
*ReleaseProfileApi* | [**ApiV3ReleaseprofileIdGet**](docs/ReleaseProfileApi.md#apiv3releaseprofileidget) | **Get** /api/v3/releaseprofile/{id} | 
*ReleaseProfileApi* | [**ApiV3ReleaseprofileIdPut**](docs/ReleaseProfileApi.md#apiv3releaseprofileidput) | **Put** /api/v3/releaseprofile/{id} | 
*ReleaseProfileApi* | [**ApiV3ReleaseprofilePost**](docs/ReleaseProfileApi.md#apiv3releaseprofilepost) | **Post** /api/v3/releaseprofile | 
*ReleasePushApi* | [**ApiV3ReleasePushPost**](docs/ReleasePushApi.md#apiv3releasepushpost) | **Post** /api/v3/release/push | 
*RemotePathMappingApi* | [**ApiV3RemotepathmappingGet**](docs/RemotePathMappingApi.md#apiv3remotepathmappingget) | **Get** /api/v3/remotepathmapping | 
*RemotePathMappingApi* | [**ApiV3RemotepathmappingIdDelete**](docs/RemotePathMappingApi.md#apiv3remotepathmappingiddelete) | **Delete** /api/v3/remotepathmapping/{id} | 
*RemotePathMappingApi* | [**ApiV3RemotepathmappingIdGet**](docs/RemotePathMappingApi.md#apiv3remotepathmappingidget) | **Get** /api/v3/remotepathmapping/{id} | 
*RemotePathMappingApi* | [**ApiV3RemotepathmappingIdPut**](docs/RemotePathMappingApi.md#apiv3remotepathmappingidput) | **Put** /api/v3/remotepathmapping/{id} | 
*RemotePathMappingApi* | [**ApiV3RemotepathmappingPost**](docs/RemotePathMappingApi.md#apiv3remotepathmappingpost) | **Post** /api/v3/remotepathmapping | 
*RenameMovieApi* | [**ApiV3RenameGet**](docs/RenameMovieApi.md#apiv3renameget) | **Get** /api/v3/rename | 
*RootFolderApi* | [**ApiV3RootfolderGet**](docs/RootFolderApi.md#apiv3rootfolderget) | **Get** /api/v3/rootfolder | 
*RootFolderApi* | [**ApiV3RootfolderIdDelete**](docs/RootFolderApi.md#apiv3rootfolderiddelete) | **Delete** /api/v3/rootfolder/{id} | 
*RootFolderApi* | [**ApiV3RootfolderIdGet**](docs/RootFolderApi.md#apiv3rootfolderidget) | **Get** /api/v3/rootfolder/{id} | 
*RootFolderApi* | [**ApiV3RootfolderPost**](docs/RootFolderApi.md#apiv3rootfolderpost) | **Post** /api/v3/rootfolder | 
*StaticResourceApi* | [**ContentPathGet**](docs/StaticResourceApi.md#contentpathget) | **Get** /content/{path} | 
*StaticResourceApi* | [**LoginGet**](docs/StaticResourceApi.md#loginget) | **Get** /login | 
*StaticResourceApi* | [**PathGet**](docs/StaticResourceApi.md#pathget) | **Get** /{path} | 
*StaticResourceApi* | [**RootGet**](docs/StaticResourceApi.md#rootget) | **Get** / | 
*SystemApi* | [**ApiV3SystemRestartPost**](docs/SystemApi.md#apiv3systemrestartpost) | **Post** /api/v3/system/restart | 
*SystemApi* | [**ApiV3SystemRoutesDuplicateGet**](docs/SystemApi.md#apiv3systemroutesduplicateget) | **Get** /api/v3/system/routes/duplicate | 
*SystemApi* | [**ApiV3SystemRoutesGet**](docs/SystemApi.md#apiv3systemroutesget) | **Get** /api/v3/system/routes | 
*SystemApi* | [**ApiV3SystemShutdownPost**](docs/SystemApi.md#apiv3systemshutdownpost) | **Post** /api/v3/system/shutdown | 
*SystemApi* | [**ApiV3SystemStatusGet**](docs/SystemApi.md#apiv3systemstatusget) | **Get** /api/v3/system/status | 
*TagApi* | [**ApiV3TagGet**](docs/TagApi.md#apiv3tagget) | **Get** /api/v3/tag | 
*TagApi* | [**ApiV3TagIdDelete**](docs/TagApi.md#apiv3tagiddelete) | **Delete** /api/v3/tag/{id} | 
*TagApi* | [**ApiV3TagIdGet**](docs/TagApi.md#apiv3tagidget) | **Get** /api/v3/tag/{id} | 
*TagApi* | [**ApiV3TagIdPut**](docs/TagApi.md#apiv3tagidput) | **Put** /api/v3/tag/{id} | 
*TagApi* | [**ApiV3TagPost**](docs/TagApi.md#apiv3tagpost) | **Post** /api/v3/tag | 
*TagDetailsApi* | [**ApiV3TagDetailGet**](docs/TagDetailsApi.md#apiv3tagdetailget) | **Get** /api/v3/tag/detail | 
*TagDetailsApi* | [**ApiV3TagDetailIdGet**](docs/TagDetailsApi.md#apiv3tagdetailidget) | **Get** /api/v3/tag/detail/{id} | 
*TaskApi* | [**ApiV3SystemTaskGet**](docs/TaskApi.md#apiv3systemtaskget) | **Get** /api/v3/system/task | 
*TaskApi* | [**ApiV3SystemTaskIdGet**](docs/TaskApi.md#apiv3systemtaskidget) | **Get** /api/v3/system/task/{id} | 
*UiConfigApi* | [**ApiV3ConfigUiGet**](docs/UiConfigApi.md#apiv3configuiget) | **Get** /api/v3/config/ui | 
*UiConfigApi* | [**ApiV3ConfigUiIdGet**](docs/UiConfigApi.md#apiv3configuiidget) | **Get** /api/v3/config/ui/{id} | 
*UiConfigApi* | [**ApiV3ConfigUiIdPut**](docs/UiConfigApi.md#apiv3configuiidput) | **Put** /api/v3/config/ui/{id} | 
*UpdateApi* | [**ApiV3UpdateGet**](docs/UpdateApi.md#apiv3updateget) | **Get** /api/v3/update | 
*UpdateLogFileApi* | [**ApiV3LogFileUpdateFilenameGet**](docs/UpdateLogFileApi.md#apiv3logfileupdatefilenameget) | **Get** /api/v3/log/file/update/{filename} | 
*UpdateLogFileApi* | [**ApiV3LogFileUpdateGet**](docs/UpdateLogFileApi.md#apiv3logfileupdateget) | **Get** /api/v3/log/file/update | 

## Documentation For Models

 - [AddMovieMethod](docs/AddMovieMethod.md)
 - [AddMovieOptions](docs/AddMovieOptions.md)
 - [AlternativeTitleResource](docs/AlternativeTitleResource.md)
 - [ApiInfoResource](docs/ApiInfoResource.md)
 - [ApplyTags](docs/ApplyTags.md)
 - [AuthenticationRequiredType](docs/AuthenticationRequiredType.md)
 - [AuthenticationType](docs/AuthenticationType.md)
 - [AutoTaggingResource](docs/AutoTaggingResource.md)
 - [AutoTaggingSpecificationSchema](docs/AutoTaggingSpecificationSchema.md)
 - [BackupResource](docs/BackupResource.md)
 - [BackupType](docs/BackupType.md)
 - [BlocklistBulkResource](docs/BlocklistBulkResource.md)
 - [BlocklistResource](docs/BlocklistResource.md)
 - [BlocklistResourcePagingResource](docs/BlocklistResourcePagingResource.md)
 - [CertificateValidationType](docs/CertificateValidationType.md)
 - [CollectionMovieResource](docs/CollectionMovieResource.md)
 - [CollectionResource](docs/CollectionResource.md)
 - [CollectionUpdateResource](docs/CollectionUpdateResource.md)
 - [ColonReplacementFormat](docs/ColonReplacementFormat.md)
 - [Command](docs/Command.md)
 - [CommandPriority](docs/CommandPriority.md)
 - [CommandResource](docs/CommandResource.md)
 - [CommandResult](docs/CommandResult.md)
 - [CommandStatus](docs/CommandStatus.md)
 - [CommandTrigger](docs/CommandTrigger.md)
 - [CreditResource](docs/CreditResource.md)
 - [CreditType](docs/CreditType.md)
 - [CustomFilterResource](docs/CustomFilterResource.md)
 - [CustomFormatBulkResource](docs/CustomFormatBulkResource.md)
 - [CustomFormatResource](docs/CustomFormatResource.md)
 - [CustomFormatSpecificationSchema](docs/CustomFormatSpecificationSchema.md)
 - [DatabaseType](docs/DatabaseType.md)
 - [DelayProfileResource](docs/DelayProfileResource.md)
 - [DiskSpaceResource](docs/DiskSpaceResource.md)
 - [DownloadClientBulkResource](docs/DownloadClientBulkResource.md)
 - [DownloadClientConfigResource](docs/DownloadClientConfigResource.md)
 - [DownloadClientResource](docs/DownloadClientResource.md)
 - [DownloadProtocol](docs/DownloadProtocol.md)
 - [ExtraFileResource](docs/ExtraFileResource.md)
 - [ExtraFileType](docs/ExtraFileType.md)
 - [Field](docs/Field.md)
 - [FileDateType](docs/FileDateType.md)
 - [HealthCheckResult](docs/HealthCheckResult.md)
 - [HealthResource](docs/HealthResource.md)
 - [HistoryResource](docs/HistoryResource.md)
 - [HistoryResourcePagingResource](docs/HistoryResourcePagingResource.md)
 - [HostConfigResource](docs/HostConfigResource.md)
 - [HttpUri](docs/HttpUri.md)
 - [ImportListBulkResource](docs/ImportListBulkResource.md)
 - [ImportListConfigResource](docs/ImportListConfigResource.md)
 - [ImportListExclusionBulkResource](docs/ImportListExclusionBulkResource.md)
 - [ImportListExclusionResource](docs/ImportListExclusionResource.md)
 - [ImportListExclusionResourcePagingResource](docs/ImportListExclusionResourcePagingResource.md)
 - [ImportListResource](docs/ImportListResource.md)
 - [ImportListType](docs/ImportListType.md)
 - [IndexerBulkResource](docs/IndexerBulkResource.md)
 - [IndexerConfigResource](docs/IndexerConfigResource.md)
 - [IndexerFlagResource](docs/IndexerFlagResource.md)
 - [IndexerResource](docs/IndexerResource.md)
 - [Language](docs/Language.md)
 - [LanguageResource](docs/LanguageResource.md)
 - [LocalizationLanguageResource](docs/LocalizationLanguageResource.md)
 - [LogFileResource](docs/LogFileResource.md)
 - [LogResource](docs/LogResource.md)
 - [LogResourcePagingResource](docs/LogResourcePagingResource.md)
 - [LoginBody](docs/LoginBody.md)
 - [ManualImportReprocessResource](docs/ManualImportReprocessResource.md)
 - [ManualImportResource](docs/ManualImportResource.md)
 - [MediaCover](docs/MediaCover.md)
 - [MediaCoverTypes](docs/MediaCoverTypes.md)
 - [MediaInfoResource](docs/MediaInfoResource.md)
 - [MediaManagementConfigResource](docs/MediaManagementConfigResource.md)
 - [MetadataConfigResource](docs/MetadataConfigResource.md)
 - [MetadataResource](docs/MetadataResource.md)
 - [Modifier](docs/Modifier.md)
 - [MonitorTypes](docs/MonitorTypes.md)
 - [MovieCollectionResource](docs/MovieCollectionResource.md)
 - [MovieEditorResource](docs/MovieEditorResource.md)
 - [MovieFileListResource](docs/MovieFileListResource.md)
 - [MovieFileResource](docs/MovieFileResource.md)
 - [MovieHistoryEventType](docs/MovieHistoryEventType.md)
 - [MovieResource](docs/MovieResource.md)
 - [MovieResourcePagingResource](docs/MovieResourcePagingResource.md)
 - [MovieRuntimeFormatType](docs/MovieRuntimeFormatType.md)
 - [MovieStatisticsResource](docs/MovieStatisticsResource.md)
 - [MovieStatusType](docs/MovieStatusType.md)
 - [NamingConfigResource](docs/NamingConfigResource.md)
 - [NotificationResource](docs/NotificationResource.md)
 - [ParseResource](docs/ParseResource.md)
 - [ParsedMovieInfo](docs/ParsedMovieInfo.md)
 - [PingResource](docs/PingResource.md)
 - [PrivacyLevel](docs/PrivacyLevel.md)
 - [ProfileFormatItemResource](docs/ProfileFormatItemResource.md)
 - [ProperDownloadTypes](docs/ProperDownloadTypes.md)
 - [ProviderMessage](docs/ProviderMessage.md)
 - [ProviderMessageType](docs/ProviderMessageType.md)
 - [ProxyType](docs/ProxyType.md)
 - [Quality](docs/Quality.md)
 - [QualityDefinitionResource](docs/QualityDefinitionResource.md)
 - [QualityModel](docs/QualityModel.md)
 - [QualityProfileQualityItemResource](docs/QualityProfileQualityItemResource.md)
 - [QualityProfileResource](docs/QualityProfileResource.md)
 - [QualitySource](docs/QualitySource.md)
 - [QueueBulkResource](docs/QueueBulkResource.md)
 - [QueueResource](docs/QueueResource.md)
 - [QueueResourcePagingResource](docs/QueueResourcePagingResource.md)
 - [QueueStatusResource](docs/QueueStatusResource.md)
 - [RatingChild](docs/RatingChild.md)
 - [RatingType](docs/RatingType.md)
 - [Ratings](docs/Ratings.md)
 - [Rejection](docs/Rejection.md)
 - [RejectionType](docs/RejectionType.md)
 - [ReleaseProfileResource](docs/ReleaseProfileResource.md)
 - [ReleaseResource](docs/ReleaseResource.md)
 - [RemotePathMappingResource](docs/RemotePathMappingResource.md)
 - [RenameMovieResource](docs/RenameMovieResource.md)
 - [RescanAfterRefreshType](docs/RescanAfterRefreshType.md)
 - [Revision](docs/Revision.md)
 - [RootFolderResource](docs/RootFolderResource.md)
 - [RuntimeMode](docs/RuntimeMode.md)
 - [SelectOption](docs/SelectOption.md)
 - [SortDirection](docs/SortDirection.md)
 - [SourceType](docs/SourceType.md)
 - [SystemResource](docs/SystemResource.md)
 - [TagDetailsResource](docs/TagDetailsResource.md)
 - [TagResource](docs/TagResource.md)
 - [TaskResource](docs/TaskResource.md)
 - [TmdbCountryCode](docs/TmdbCountryCode.md)
 - [TrackedDownloadState](docs/TrackedDownloadState.md)
 - [TrackedDownloadStatus](docs/TrackedDownloadStatus.md)
 - [TrackedDownloadStatusMessage](docs/TrackedDownloadStatusMessage.md)
 - [UiConfigResource](docs/UiConfigResource.md)
 - [UnmappedFolder](docs/UnmappedFolder.md)
 - [UpdateChanges](docs/UpdateChanges.md)
 - [UpdateMechanism](docs/UpdateMechanism.md)
 - [UpdateResource](docs/UpdateResource.md)

## Documentation For Authorization

## X-Api-Key
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```
## apikey
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author

