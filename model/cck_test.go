package model

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"testing"
)

const (
	INSERT_INTO_CONTENT_NODE_FIELD          = `INSERT INTO content_node_field (field_name, type, global_settings, required, multiple, db_storage, module, db_columns, active, locked) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	INSERT_INTO_CONTENT_NODE_FIELD_INSTANCE = `INSERT INTO content_node_field_instance ( field_name, type_name, weight, label, widget_type, widget_settings, display_settings, description, widget_module, widget_active ) values ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )`
	INSERT_INTO_CONTENT_TYPE_BLOG           = `INSERT INTO content_type_blog (vid, nid, field_video_embed, field_video_value, field_video_provider, field_video_data, field_video_ref_embed, field_video_ref_value, field_video_ref_provider, field_video_ref_data, field_video_ref_version, field_video_ref_duration) values ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )`
)

func TestGetCCKFields(t *testing.T) {
	os.Remove("./test.db")
	defer os.Remove("./test.db")

	db := Connect("sqlite3", "./test.db", "", true)

	_, err := db.Db.Exec(`CREATE TABLE IF NOT EXISTS content_node_field (
	  field_name varchar(32) NOT NULL DEFAULT '',
	  type varchar(127) NOT NULL DEFAULT '',
	  global_settings TEXT NOT NULL,
	  required int NOT NULL DEFAULT '0',
	  multiple int NOT NULL DEFAULT '0',
	  db_storage int NOT NULL DEFAULT '0',
	  module varchar(127) NOT NULL DEFAULT '',
	  db_columns TEXT NOT NULL,
	  active tinyint NOT NULL DEFAULT '0',
	  locked tinyint NOT NULL DEFAULT '0',
	  PRIMARY KEY (field_name)
	)`)
	if err != nil {
		t.Error(err)
	}
	_, err = db.Db.Exec(`CREATE TABLE IF NOT EXISTS content_node_field_instance (
	  field_name varchar(32) NOT NULL DEFAULT '',
	  type_name varchar(32) NOT NULL DEFAULT '',
	  weight int NOT NULL DEFAULT '0',
	  label varchar(255) NOT NULL DEFAULT '',
	  widget_type varchar(32) NOT NULL DEFAULT '',
	  widget_settings TEXT NOT NULL,
	  display_settings TEXT NOT NULL,
	  description TEXT NOT NULL,
	  widget_module varchar(127) NOT NULL DEFAULT '',
	  widget_active tinyint NOT NULL DEFAULT '0',
	  PRIMARY KEY (field_name,type_name)
	)`)
	if err != nil {
		t.Error(err)
	}
	_, err = db.Db.Exec(`CREATE TABLE IF NOT EXISTS content_type_blog (
	  vid int NOT NULL DEFAULT '0',
	  nid int NOT NULL DEFAULT '0',
	  field_video_embed TEXT,
	  field_video_value varchar(255) NOT NULL DEFAULT '',
	  field_video_provider varchar(255) NOT NULL DEFAULT '',
	  field_video_data TEXT,
	  field_video_ref_embed TEXT,
	  field_video_ref_value varchar(255) DEFAULT NULL,
	  field_video_ref_provider varchar(255) DEFAULT NULL,
	  field_video_ref_data TEXT,
	  field_video_ref_version int unsigned NOT NULL DEFAULT '0',
	  field_video_ref_duration int unsigned NOT NULL DEFAULT '0',
	  PRIMARY KEY (vid)
	)`)
	if err != nil {
		t.Error(err)
	}

	insert(t, db.Db, INSERT_INTO_CONTENT_NODE_FIELD, "field_video", "video_cck", "a:0:{}", 0, 0, 1, "", "", 0, 0)
	insert(t, db.Db, INSERT_INTO_CONTENT_NODE_FIELD, "field_video_ref", "emvideo", "a:0:{}", 0, 0, 1, "emvideo", `a:6:{s:5:"embed";a:4:{s:4:"type";s:4:"text";s:4:"size";s:3:"big";s:8:"not null";b:0;s:8:"sortable";b:1;}s:5:"value";a:4:{s:4:"type";s:7:"varchar";s:6:"length";i:255;s:8:"not null";b:0;s:8:"sortable";b:1;}s:8:"provider";a:4:{s:4:"type";s:7:"varchar";s:6:"length";i:255;s:8:"not null";b:0;s:8:"sortable";b:1;}s:4:"data";a:4:{s:4:"type";s:4:"text";s:4:"size";s:3:"big";s:8:"not null";b:0;s:8:"sortable";b:0;}s:7:"version";a:5:{s:11:"description";s:35:"The version of the provider\'s data.";s:4:"type";s:3:"int";s:8:"unsigned";b:1;s:8:"not null";b:1;s:7:"default";i:0;}s:8:"duration";a:5:{s:11:"description";s:41:"Store the duration of a video in seconds.";s:4:"type";s:3:"int";s:8:"unsigned";b:1;s:8:"not null";b:1;s:7:"default";i:0;}}`, 1, 0)
	insert(t, db.Db, INSERT_INTO_CONTENT_NODE_FIELD_INSTANCE, "field_video", "blog", -2, "Video", "video_cck_textfields", `a:20:{s:13:"default_value";a:1:{i:0;a:2:{s:5:"embed";s:0:"";s:5:"value";s:0:"";}}s:17:"default_value_php";s:0:"";s:11:"video_width";s:3:"425";s:12:"video_height";s:3:"350";s:14:"video_autoplay";i:0;s:13:"preview_width";s:3:"425";s:14:"preview_height";s:3:"350";s:16:"preview_autoplay";i:0;s:15:"thumbnail_width";s:3:"120";s:16:"thumbnail_height";s:2:"90";s:22:"thumbnail_default_path";s:0:"";s:9:"providers";a:19:{s:6:"bliptv";i:0;s:10:"brightcove";i:0;s:11:"dailymotion";i:0;s:6:"google";i:0;s:4:"guba";i:0;s:5:"imeem";i:0;s:7:"jumpcut";i:0;s:6:"lastfm";i:0;s:9:"livevideo";i:0;s:8:"metacafe";i:0;s:7:"myspace";i:0;s:6:"revver";i:0;s:9:"sevenload";i:0;s:5:"spike";i:0;s:5:"tudou";i:0;s:4:"veoh";i:0;s:5:"vimeo";i:0;s:7:"youtube";i:0;s:14:"zzz_custom_url";i:0;}s:8:"emimport";i:1;s:7:"emthumb";i:1;s:13:"emthumb_label";s:22:"Video custom thumbnail";s:19:"emthumb_description";s:158:"If you upload a custom thumbnail, then this will be displayed when the video thumbnail is called for, overriding any automatic thumbnails by custom providers.";s:22:"emthumb_max_resolution";s:1:"0";s:19:"emimport_image_path";s:13:"videos/thumbs";s:18:"emthumb_custom_alt";i:1;s:20:"emthumb_custom_title";i:1;}`, `a:3:{s:5:"label";a:1:{s:6:"format";s:5:"above";}s:4:"full";a:2:{s:6:"format";s:7:"default";s:7:"exclude";i:0;}s:6:"teaser";a:2:{s:6:"format";s:7:"default";s:7:"exclude";i:0;}}`, "", "", 0)
	insert(t, db.Db, INSERT_INTO_CONTENT_NODE_FIELD_INSTANCE, "field_video_ref", "blog", -1, "Video", "emvideo_textfields", `a:21:{s:11:"video_width";s:3:"425";s:12:"video_height";s:3:"350";s:14:"video_autoplay";i:0;s:13:"preview_width";s:3:"425";s:14:"preview_height";s:3:"350";s:16:"preview_autoplay";i:0;s:15:"thumbnail_width";s:3:"120";s:16:"thumbnail_height";s:2:"90";s:22:"thumbnail_default_path";s:0:"";s:20:"thumbnail_link_title";s:9:"See video";s:9:"providers";a:22:{s:7:"archive";s:7:"archive";s:6:"bliptv";s:6:"bliptv";s:10:"brightcove";s:10:"brightcove";s:11:"dailymotion";s:11:"dailymotion";s:6:"google";s:6:"google";s:4:"guba";s:4:"guba";s:5:"imeem";s:5:"imeem";s:7:"jumpcut";s:7:"jumpcut";s:6:"lastfm";s:6:"lastfm";s:9:"livevideo";s:9:"livevideo";s:8:"metacafe";s:8:"metacafe";s:7:"myspace";s:7:"myspace";s:6:"revver";s:6:"revver";s:9:"sevenload";s:9:"sevenload";s:5:"spike";s:5:"spike";s:5:"tudou";s:5:"tudou";s:7:"ustream";s:7:"ustream";s:11:"ustreamlive";s:11:"ustreamlive";s:4:"veoh";s:4:"veoh";s:5:"vimeo";s:5:"vimeo";s:7:"youtube";s:7:"youtube";s:14:"zzz_custom_url";s:14:"zzz_custom_url";}s:8:"emimport";i:1;s:7:"emthumb";i:1;s:13:"emthumb_label";s:22:"Video custom thumbnail";s:19:"emthumb_description";s:158:"If you upload a custom thumbnail, then this will be displayed when the Video thumbnail is called for, overriding any automatic thumbnails by custom providers.";s:22:"emthumb_max_resolution";s:1:"0";s:19:"emimport_image_path";s:0:"";s:18:"emthumb_custom_alt";i:0;s:20:"emthumb_custom_title";i:0;s:13:"default_value";a:1:{i:0;a:3:{s:5:"embed";s:0:"";s:5:"value";s:0:"";s:7:"emthumb";a:1:{s:7:"emthumb";a:1:{s:7:"emthumb";s:0:"";}}}}s:17:"default_value_php";N;}`, `a:8:{s:6:"weight";i:0;s:6:"parent";s:0:"";s:5:"label";a:1:{s:6:"format";s:5:"above";}s:6:"teaser";a:2:{s:6:"format";s:7:"default";s:7:"exclude";i:0;}s:4:"full";a:2:{s:6:"format";s:7:"default";s:7:"exclude";i:0;}i:4;a:2:{s:6:"format";s:7:"default";s:7:"exclude";i:0;}i:2;a:2:{s:6:"format";s:7:"default";s:7:"exclude";i:0;}i:3;a:2:{s:6:"format";s:7:"default";s:7:"exclude";i:0;}}`, "", "emvideo", 1)

	insert(t, db.Db, INSERT_INTO_CONTENT_TYPE_BLOG, 32, 32, "<object width=\"425\" height=\"350\"><param name=\"movie\" value=\"http://youtube.com/v/xz3v1TnCvVA\"></param><embed src=\"http://youtube.com/v/xz3v1TnCvVA\" type=\"application/x-shockwave-flash\" width=\"425\" height=\"350\"></embed></object>", "xz3v1TnCvVA", "youtube", "a:4:{s:25:\"video_cck_youtube_version\";i:1;s:9:\"thumbnail\";a:1:{s:3:\"url\";s:43:\"http://img.youtube.com/vi/xz3v1TnCvVA/0.jpg\";}s:5:\"flash\";a:3:{s:3:\"url\";s:32:\"http://youtube.com/v/xz3v1TnCvVA\";s:4:\"size\";s:3:\"817\";s:4:\"mime\";s:29:\"application/x-shockwave-flash\";}s:7:\"emthumb\";a:0:{}}", "http://youtube.com/v/xz3v1TnCvVA", "xz3v1TnCvVA", "youtube", "a:7:{s:20:\"emvideo_data_version\";i:3;s:23:\"emvideo_youtube_version\";i:3;s:3:\"raw\";a:0:{}s:8:\"duration\";i:0;s:8:\"playlist\";i:0;s:9:\"thumbnail\";a:1:{s:3:\"url\";s:43:\"http://img.youtube.com/vi/xz3v1TnCvVA/0.jpg\";}s:5:\"flash\";a:3:{s:3:\"url\";s:32:\"http://youtube.com/v/xz3v1TnCvVA\";s:4:\"size\";s:4:\"1178\";s:4:\"mime\";s:29:\"application/x-shockwave-flash\";}}", 3, 0)
	insert(t, db.Db, INSERT_INTO_CONTENT_TYPE_BLOG, 33, 33, "http://youtube.com/watch?v=dU4mjRCxK1Y", "dU4mjRCxK1Y", "youtube", "a:4:{s:25:\"video_cck_youtube_version\";i:1;s:9:\"thumbnail\";a:1:{s:3:\"url\";s:43:\"http://img.youtube.com/vi/dU4mjRCxK1Y/0.jpg\";}s:5:\"flash\";a:3:{s:3:\"url\";s:32:\"http://youtube.com/v/dU4mjRCxK1Y\";s:4:\"size\";s:3:\"817\";s:4:\"mime\";s:29:\"application/x-shockwave-flash\";}s:7:\"emthumb\";a:0:{}}", "http://youtube.com/watch?v=dU4mjRCxK1Y", "dU4mjRCxK1Y", "youtube", "a:7:{s:20:\"emvideo_data_version\";i:3;s:23:\"emvideo_youtube_version\";i:3;s:3:\"raw\";a:6:{s:16:\"_emfield_arghash\";s:25:\"youtube:videodU4mjRCxK1Y:\";s:27:\"http://www.w3.org/2005/Atom\";a:9:{s:2:\"ID\";a:2:{i:0;s:53:\"http://gdata.youtube.com/feeds/api/videos/dU4mjRCxK1Y\";i:1;N;}s:9:\"PUBLISHED\";a:2:{i:0;s:24:\"2008-04-06T13:07:09.000Z\";i:1;N;}s:7:\"UPDATED\";a:2:{i:0;s:24:\"2012-09-24T09:12:24.000Z\";i:1;N;}s:8:\"CATEGORY\";a:4:{i:0;N;i:1;a:2:{s:6:\"SCHEME\";s:37:\"http://schemas.google.com/g/2005#kind\";s:4:\"TERM\";s:43:\"http://gdata.youtube.com/schemas/2007#video\";}i:2;N;i:3;a:3:{s:6:\"SCHEME\";s:52:\"http://gdata.youtube.com/schemas/2007/categories.cat\";s:4:\"TERM\";s:9:\"Education\";s:5:\"LABEL\";s:9:\"Education\";}}s:5:\"TITLE\";a:2:{i:0;s:26:\"\"Coming Out\" as an Atheist\";i:1;a:1:{s:4:\"TYPE\";s:4:\"text\";}}s:7:\"CONTENT\";a:2:{i:0;s:585:\"I use the phrase \"Coming Out\" because in many ways coming out as an Atheist is very similar to coming out as gay.   It really depends on the types of people that make up your family but I am sure there are some people for whom coming out as an atheist would actually be harder then coming out as gay.  I seem to remember someone making a joke about telling your parents you are an atheist so that you can say that you were just kidding and that you were just gay. \n\nSee this video for the reaction that young ex-Catholic got when he came out: http://www.youtube.com/watch?v=P8Aq00yJSxo\";i:1;a:1:{s:4:\"TYPE\";s:4:\"text\";}}s:4:\"LINK\";a:12:{i:0;N;i:1;a:3:{s:3:\"REL\";s:58:\"http://gdata.youtube.com/schemas/2007#video.in-response-to\";s:4:\"TYPE\";s:20:\"application/atom+xml\";s:4:\"HREF\";s:53:\"http://gdata.youtube.com/feeds/api/videos/5OXoveg5VRw\";}i:2;N;i:3;a:3:{s:3:\"REL\";s:9:\"alternate\";s:4:\"TYPE\";s:9:\"text/html\";s:4:\"HREF\";s:64:\"http://www.youtube.com/watch?v=dU4mjRCxK1Y&feature=youtube_gdata\";}i:4;N;i:5;a:3:{s:3:\"REL\";s:53:\"http://gdata.youtube.com/schemas/2007#video.responses\";s:4:\"TYPE\";s:20:\"application/atom+xml\";s:4:\"HREF\";s:63:\"http://gdata.youtube.com/feeds/api/videos/dU4mjRCxK1Y/responses\";}i:6;N;i:7;a:3:{s:3:\"REL\";s:51:\"http://gdata.youtube.com/schemas/2007#video.related\";s:4:\"TYPE\";s:20:\"application/atom+xml\";s:4:\"HREF\";s:61:\"http://gdata.youtube.com/feeds/api/videos/dU4mjRCxK1Y/related\";}i:8;N;i:9;a:3:{s:3:\"REL\";s:44:\"http://gdata.youtube.com/schemas/2007#mobile\";s:4:\"TYPE\";s:9:\"text/html\";s:4:\"HREF\";s:42:\"http://m.youtube.com/details?v=dU4mjRCxK1Y\";}i:10;N;i:11;a:3:{s:3:\"REL\";s:4:\"self\";s:4:\"TYPE\";s:20:\"application/atom+xml\";s:4:\"HREF\";s:53:\"http://gdata.youtube.com/feeds/api/videos/dU4mjRCxK1Y\";}}s:6:\"AUTHOR\";a:1:{s:4:\"NAME\";a:2:{i:0;s:14:\"atheistliberty\";i:1;N;}}s:13:\"YT:STATISTICS\";a:2:{i:0;N;i:1;a:2:{s:13:\"FAVORITECOUNT\";s:1:\"0\";s:9:\"VIEWCOUNT\";s:4:\"6090\";}}}s:6:\"AUTHOR\";a:1:{s:3:\"URI\";a:2:{i:0;s:55:\"http://gdata.youtube.com/feeds/api/users/atheistliberty\";i:1;N;}}s:11:\"GD:COMMENTS\";a:1:{s:11:\"GD:FEEDLINK\";a:2:{i:0;N;i:1;a:3:{s:3:\"REL\";s:46:\"http://gdata.youtube.com/schemas/2007#comments\";s:4:\"HREF\";s:62:\"http://gdata.youtube.com/feeds/api/videos/dU4mjRCxK1Y/comments\";s:9:\"COUNTHINT\";s:3:\"155\";}}}s:11:\"MEDIA:GROUP\";a:8:{s:14:\"MEDIA:CATEGORY\";a:2:{i:0;s:9:\"Education\";i:1;a:2:{s:5:\"LABEL\";s:9:\"Education\";s:6:\"SCHEME\";s:52:\"http://gdata.youtube.com/schemas/2007/categories.cat\";}}s:13:\"MEDIA:CONTENT\";a:6:{i:0;N;i:1;a:7:{s:3:\"URL\";s:73:\"http://www.youtube.com/v/dU4mjRCxK1Y?version=3&f=videos&app=youtube_gdata\";s:4:\"TYPE\";s:29:\"application/x-shockwave-flash\";s:6:\"MEDIUM\";s:5:\"video\";s:9:\"ISDEFAULT\";s:4:\"true\";s:10:\"EXPRESSION\";s:4:\"full\";s:8:\"DURATION\";s:3:\"658\";s:9:\"YT:FORMAT\";s:1:\"5\";}i:2;N;i:3;a:6:{s:3:\"URL\";s:95:\"rtsp://v7.cache5.c.youtube.com/CiILENy73wIaGQlWK7EQjSZOdRMYDSANFEgGUgZ2aWRlb3MM/0/0/0/video.3gp\";s:4:\"TYPE\";s:10:\"video/3gpp\";s:6:\"MEDIUM\";s:5:\"video\";s:10:\"EXPRESSION\";s:4:\"full\";s:8:\"DURATION\";s:3:\"658\";s:9:\"YT:FORMAT\";s:1:\"1\";}i:4;N;i:5;a:6:{s:3:\"URL\";s:95:\"rtsp://v8.cache6.c.youtube.com/CiILENy73wIaGQlWK7EQjSZOdRMYESARFEgGUgZ2aWRlb3MM/0/0/0/video.3gp\";s:4:\"TYPE\";s:10:\"video/3gpp\";s:6:\"MEDIUM\";s:5:\"video\";s:10:\"EXPRESSION\";s:4:\"full\";s:8:\"DURATION\";s:3:\"658\";s:9:\"YT:FORMAT\";s:1:\"6\";}}s:17:\"MEDIA:DESCRIPTION\";a:2:{i:0;s:585:\"I use the phrase \"Coming Out\" because in many ways coming out as an Atheist is very similar to coming out as gay.   It really depends on the types of people that make up your family but I am sure there are some people for whom coming out as an atheist would actually be harder then coming out as gay.  I seem to remember someone making a joke about telling your parents you are an atheist so that you can say that you were just kidding and that you were just gay. \n\nSee this video for the reaction that young ex-Catholic got when he came out: http://www.youtube.com/watch?v=P8Aq00yJSxo\";i:1;a:1:{s:4:\"TYPE\";s:5:\"plain\";}}s:14:\"MEDIA:KEYWORDS\";a:2:{i:0;N;i:1;N;}s:12:\"MEDIA:PLAYER\";a:2:{i:0;N;i:1;a:1:{s:3:\"URL\";s:71:\"http://www.youtube.com/watch?v=dU4mjRCxK1Y&feature=youtube_gdata_player\";}}s:15:\"MEDIA:THUMBNAIL\";a:8:{i:0;N;i:1;a:4:{s:3:\"URL\";s:39:\"http://i.ytimg.com/vi/dU4mjRCxK1Y/0.jpg\";s:6:\"HEIGHT\";s:3:\"360\";s:5:\"WIDTH\";s:3:\"480\";s:4:\"TIME\";s:8:\"00:05:29\";}i:2;N;i:3;a:4:{s:3:\"URL\";s:39:\"http://i.ytimg.com/vi/dU4mjRCxK1Y/1.jpg\";s:6:\"HEIGHT\";s:2:\"90\";s:5:\"WIDTH\";s:3:\"120\";s:4:\"TIME\";s:12:\"00:02:44.500\";}i:4;N;i:5;a:4:{s:3:\"URL\";s:39:\"http://i.ytimg.com/vi/dU4mjRCxK1Y/2.jpg\";s:6:\"HEIGHT\";s:2:\"90\";s:5:\"WIDTH\";s:3:\"120\";s:4:\"TIME\";s:8:\"00:05:29\";}i:6;N;i:7;a:4:{s:3:\"URL\";s:39:\"http://i.ytimg.com/vi/dU4mjRCxK1Y/3.jpg\";s:6:\"HEIGHT\";s:2:\"90\";s:5:\"WIDTH\";s:3:\"120\";s:4:\"TIME\";s:12:\"00:08:13.500\";}}s:11:\"MEDIA:TITLE\";a:2:{i:0;s:26:\"\"Coming Out\" as an Atheist\";i:1;a:1:{s:4:\"TYPE\";s:5:\"plain\";}}s:11:\"YT:DURATION\";a:2:{i:0;N;i:1;a:1:{s:7:\"SECONDS\";s:3:\"658\";}}}s:9:\"GD:RATING\";a:2:{i:0;N;i:1;a:5:{s:7:\"AVERAGE\";s:8:\"4.766234\";s:3:\"MAX\";s:1:\"5\";s:3:\"MIN\";s:1:\"1\";s:9:\"NUMRATERS\";s:3:\"154\";s:3:\"REL\";s:40:\"http://schemas.google.com/g/2005#overall\";}}}s:8:\"duration\";i:658;s:8:\"playlist\";i:0;s:9:\"thumbnail\";a:1:{s:3:\"url\";s:43:\"http://img.youtube.com/vi/dU4mjRCxK1Y/0.jpg\";}s:5:\"flash\";a:3:{s:3:\"url\";s:32:\"http://youtube.com/v/dU4mjRCxK1Y\";s:4:\"size\";s:4:\"1175\";s:4:\"mime\";s:29:\"application/x-shockwave-flash\";}}", 3, 658)
	insert(t, db.Db, INSERT_INTO_CONTENT_TYPE_BLOG, 34, 34, "http://youtube.com/watch?v=Jd3zEg18X48",
		"Jd3zEg18X48",
		"youtube",
		"a:4:{s:25:\"video_cck_youtube_version\";i:1;s:9:\"thumbnail\";a:1:{s:3:\"url\";s:43:\"http://img.youtube.com/vi/Jd3zEg18X48/0.jpg\";}s:5:\"flash\";a:3:{s:3:\"url\";s:32:\"http://youtube.com/v/Jd3zEg18X48\";s:4:\"size\";s:3:\"817\";s:4:\"mime\";s:29:\"application/x-shockwave-flash\";}s:7:\"emthumb\";a:0:{}}",
		"http://youtube.com/watch?v=Jd3zEg18X48",
		"Jd3zEg18X48",
		"youtube",
		"a:7:{s:20:\"emvideo_data_version\";i:3;s:23:\"emvideo_youtube_version\";i:3;s:3:\"raw\";a:6:{s:16:\"_emfield_arghash\";s:25:\"youtube:videoJd3zEg18X48:\";s:27:\"http://www.w3.org/2005/Atom\";a:9:{s:2:\"ID\";a:2:{i:0;s:53:\"http://gdata.youtube.com/feeds/api/videos/Jd3zEg18X48\";i:1;N;}s:9:\"PUBLISHED\";a:2:{i:0;s:24:\"2008-05-02T20:16:24.000Z\";i:1;N;}s:7:\"UPDATED\";a:2:{i:0;s:24:\"2012-04-15T16:24:32.000Z\";i:1;N;}s:8:\"CATEGORY\";a:4:{i:0;N;i:1;a:2:{s:6:\"SCHEME\";s:37:\"http://schemas.google.com/g/2005#kind\";s:4:\"TERM\";s:43:\"http://gdata.youtube.com/schemas/2007#video\";}i:2;N;i:3;a:3:{s:6:\"SCHEME\";s:52:\"http://gdata.youtube.com/schemas/2007/categories.cat\";s:4:\"TERM\";s:9:\"Education\";s:5:\"LABEL\";s:9:\"Education\";}}s:5:\"TITLE\";a:2:{i:0;s:28:\"Expelled: No Mystery Allowed\";i:1;a:1:{s:4:\"TYPE\";s:4:\"text\";}}s:7:\"CONTENT\";a:2:{i:0;s:688:\"http://atheistliberty.com/node/34\n\nHere is a link to one of Thunderf00t's debunkings of Ben Stein's Ravings:\nhttp://www.youtube.com/watch?v=ihYq2dGa29M\n\nExpelled Trailer and other:\nhttp://www.youtube.com/watch?v=JEPqLKErXpI\n\nHere is a video about how PZ Myers was Expelled by the film's producers from a screening of the film while his guest, Richard Dawkins, was not.  Very funny story.\nhttp://www.youtube.com/watch?v=c39jYgsvUOY\n\nHere's a link to PZ Myers' Blog:\nhttp://scienceblogs.com/pharyngula/\n\nAnd a link to Richard Dawkin's website:\nhttp://www.richarddawkins.net\n\nAlso, check out this site.  It very thoroughly debunks the claims made in the movie:\nhttp://www.expelledexposed.com\";i:1;a:1:{s:4:\"TYPE\";s:4:\"text\";}}s:4:\"LINK\";a:8:{i:0;N;i:1;a:3:{s:3:\"REL\";s:9:\"alternate\";s:4:\"TYPE\";s:9:\"text/html\";s:4:\"HREF\";s:64:\"http://www.youtube.com/watch?v=Jd3zEg18X48&feature=youtube_gdata\";}i:2;N;i:3;a:3:{s:3:\"REL\";s:53:\"http://gdata.youtube.com/schemas/2007#video.responses\";s:4:\"TYPE\";s:20:\"application/atom+xml\";s:4:\"HREF\";s:63:\"http://gdata.youtube.com/feeds/api/videos/Jd3zEg18X48/responses\";}i:4;N;i:5;a:3:{s:3:\"REL\";s:51:\"http://gdata.youtube.com/schemas/2007#video.related\";s:4:\"TYPE\";s:20:\"application/atom+xml\";s:4:\"HREF\";s:61:\"http://gdata.youtube.com/feeds/api/videos/Jd3zEg18X48/related\";}i:6;N;i:7;a:3:{s:3:\"REL\";s:4:\"self\";s:4:\"TYPE\";s:20:\"application/atom+xml\";s:4:\"HREF\";s:53:\"http://gdata.youtube.com/feeds/api/videos/Jd3zEg18X48\";}}s:6:\"AUTHOR\";a:1:{s:4:\"NAME\";a:2:{i:0;s:14:\"atheistliberty\";i:1;N;}}s:13:\"YT:STATISTICS\";a:2:{i:0;N;i:1;a:2:{s:13:\"FAVORITECOUNT\";s:1:\"0\";s:9:\"VIEWCOUNT\";s:4:\"1950\";}}}s:6:\"AUTHOR\";a:1:{s:3:\"URI\";a:2:{i:0;s:55:\"http://gdata.youtube.com/feeds/api/users/atheistliberty\";i:1;N;}}s:11:\"GD:COMMENTS\";a:1:{s:11:\"GD:FEEDLINK\";a:2:{i:0;N;i:1;a:3:{s:3:\"REL\";s:46:\"http://gdata.youtube.com/schemas/2007#comments\";s:4:\"HREF\";s:62:\"http://gdata.youtube.com/feeds/api/videos/Jd3zEg18X48/comments\";s:9:\"COUNTHINT\";s:3:\"161\";}}}s:11:\"MEDIA:GROUP\";a:8:{s:14:\"MEDIA:CATEGORY\";a:2:{i:0;s:9:\"Education\";i:1;a:2:{s:5:\"LABEL\";s:9:\"Education\";s:6:\"SCHEME\";s:52:\"http://gdata.youtube.com/schemas/2007/categories.cat\";}}s:13:\"MEDIA:CONTENT\";a:6:{i:0;N;i:1;a:7:{s:3:\"URL\";s:73:\"http://www.youtube.com/v/Jd3zEg18X48?version=3&f=videos&app=youtube_gdata\";s:4:\"TYPE\";s:29:\"application/x-shockwave-flash\";s:6:\"MEDIUM\";s:5:\"video\";s:9:\"ISDEFAULT\";s:4:\"true\";s:10:\"EXPRESSION\";s:4:\"full\";s:8:\"DURATION\";s:3:\"526\";s:9:\"YT:FORMAT\";s:1:\"5\";}i:2;N;i:3;a:6:{s:3:\"URL\";s:95:\"rtsp://v8.cache3.c.youtube.com/CiILENy73wIaGQmPX3wNEvPdJRMYDSANFEgGUgZ2aWRlb3MM/0/0/0/video.3gp\";s:4:\"TYPE\";s:10:\"video/3gpp\";s:6:\"MEDIUM\";s:5:\"video\";s:10:\"EXPRESSION\";s:4:\"full\";s:8:\"DURATION\";s:3:\"526\";s:9:\"YT:FORMAT\";s:1:\"1\";}i:4;N;i:5;a:6:{s:3:\"URL\";s:95:\"rtsp://v8.cache4.c.youtube.com/CiILENy73wIaGQmPX3wNEvPdJRMYESARFEgGUgZ2aWRlb3MM/0/0/0/video.3gp\";s:4:\"TYPE\";s:10:\"video/3gpp\";s:6:\"MEDIUM\";s:5:\"video\";s:10:\"EXPRESSION\";s:4:\"full\";s:8:\"DURATION\";s:3:\"526\";s:9:\"YT:FORMAT\";s:1:\"6\";}}s:17:\"MEDIA:DESCRIPTION\";a:2:{i:0;s:688:\"http://atheistliberty.com/node/34\n\nHere is a link to one of Thunderf00t's debunkings of Ben Stein's Ravings:\nhttp://www.youtube.com/watch?v=ihYq2dGa29M\n\nExpelled Trailer and other:\nhttp://www.youtube.com/watch?v=JEPqLKErXpI\n\nHere is a video about how PZ Myers was Expelled by the film's producers from a screening of the film while his guest, Richard Dawkins, was not.  Very funny story.\nhttp://www.youtube.com/watch?v=c39jYgsvUOY\n\nHere's a link to PZ Myers' Blog:\nhttp://scienceblogs.com/pharyngula/\n\nAnd a link to Richard Dawkin's website:\nhttp://www.richarddawkins.net\n\nAlso, check out this site.  It very thoroughly debunks the claims made in the movie:\nhttp://www.expelledexposed.com\";i:1;a:1:{s:4:\"TYPE\";s:5:\"plain\";}}s:14:\"MEDIA:KEYWORDS\";a:2:{i:0;N;i:1;N;}s:12:\"MEDIA:PLAYER\";a:2:{i:0;N;i:1;a:1:{s:3:\"URL\";s:71:\"http://www.youtube.com/watch?v=Jd3zEg18X48&feature=youtube_gdata_player\";}}s:15:\"MEDIA:THUMBNAIL\";a:8:{i:0;N;i:1;a:4:{s:3:\"URL\";s:39:\"http://i.ytimg.com/vi/Jd3zEg18X48/0.jpg\";s:6:\"HEIGHT\";s:3:\"360\";s:5:\"WIDTH\";s:3:\"480\";s:4:\"TIME\";s:8:\"00:04:23\";}i:2;N;i:3;a:4:{s:3:\"URL\";s:39:\"http://i.ytimg.com/vi/Jd3zEg18X48/1.jpg\";s:6:\"HEIGHT\";s:2:\"90\";s:5:\"WIDTH\";s:3:\"120\";s:4:\"TIME\";s:12:\"00:02:11.500\";}i:4;N;i:5;a:4:{s:3:\"URL\";s:39:\"http://i.ytimg.com/vi/Jd3zEg18X48/2.jpg\";s:6:\"HEIGHT\";s:2:\"90\";s:5:\"WIDTH\";s:3:\"120\";s:4:\"TIME\";s:8:\"00:04:23\";}i:6;N;i:7;a:4:{s:3:\"URL\";s:39:\"http://i.ytimg.com/vi/Jd3zEg18X48/3.jpg\";s:6:\"HEIGHT\";s:2:\"90\";s:5:\"WIDTH\";s:3:\"120\";s:4:\"TIME\";s:12:\"00:06:34.500\";}}s:11:\"MEDIA:TITLE\";a:2:{i:0;s:28:\"Expelled: No Mystery Allowed\";i:1;a:1:{s:4:\"TYPE\";s:5:\"plain\";}}s:11:\"YT:DURATION\";a:2:{i:0;N;i:1;a:1:{s:7:\"SECONDS\";s:3:\"526\";}}}s:9:\"GD:RATING\";a:2:{i:0;N;i:1;a:5:{s:7:\"AVERAGE\";s:9:\"4.1641793\";s:3:\"MAX\";s:1:\"5\";s:3:\"MIN\";s:1:\"1\";s:9:\"NUMRATERS\";s:2:\"67\";s:3:\"REL\";s:40:\"http://schemas.google.com/g/2005#overall\";}}}s:8:\"duration\";i:526;s:8:\"playlist\";i:0;s:9:\"thumbnail\";a:1:{s:3:\"url\";s:43:\"http://img.youtube.com/vi/Jd3zEg18X48/0.jpg\";}s:5:\"flash\";a:3:{s:3:\"url\";s:32:\"http://youtube.com/v/Jd3zEg18X48\";s:4:\"size\";s:4:\"1178\";s:4:\"mime\";s:29:\"application/x-shockwave-flash\";}}",
		3,
		526)

	fields, err := db.CCKFields()
	if err != nil {
		t.Error(err)
	}

	blog, ok := fields["blog"]
	if !ok {
		t.Error("no CCKFields for blog node type!")
	}
	if len(blog) != 1 {
		t.Error("too many CCK Fields returned: %d", len(blog))
	}
	if blog[0].Name != "field_video_ref" {
		t.Error("unexpected field returned: %s", blog[0].Name)
	}

	fieldData, err := db.CCKDataForNode(&Node{Nid: 32, Vid: 32, Type: "blog"}, fields["blog"])
	if err != nil {
		t.Error("Getting field data", err)
	}
	fieldId := CCKField{"field_video_ref", "value", "varchar"}
	if fieldData[fieldId] != "xz3v1TnCvVA" {
		t.Errorf("field_video_ref_value did not contain the expected value: %T(%s)", fieldData[fieldId], fieldData[fieldId])
	}
	for k, v := range fieldData {
		switch value := v.(type) {
		case string:
			fmt.Printf("%s = %s\n", k, value)
		case fmt.Stringer:
			fmt.Printf("%s = %s\n", k, value.String())
		case int64:
			fmt.Printf("%s = %d\n", k, value)
		default:
			fmt.Printf("%s = %s\n", k, value)
		}
	}
}

func insert(t *testing.T, db *sql.DB, insert string, args ...interface{}) error {
	r, err := db.Exec(insert, args...)
	if err != nil {
		t.Error("error executing insert: ", err)
	}
	ra, err := r.RowsAffected()
	if err != nil {
		t.Error("error retrieving rows affected: ", err)
	}
	if ra != 1 {
		t.Error("Incorrect number of rows affected on insert: ", r)
	}
	return err
}
