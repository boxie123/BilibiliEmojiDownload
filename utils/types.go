package utils

type Emoji struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    EmojiData `json:"data"`
	TTL     int       `json:"ttl"`
}

type EmojiMeta struct {
	Size       int    `json:"size"`
	LabelText  string `json:"label_text"`
	LabelColor string `json:"label_color"`
	ItemID     int64  `json:"item_id"`
	ItemURL    string `json:"item_url"`
	AssetID    int    `json:"asset_id"`
}

type EmojiMeta0 struct {
	Size  int    `json:"size"`
	Alias string `json:"alias"`
}

type EmojiEmotes struct {
	ID        int        `json:"id"`
	PackageID int        `json:"package_id"`
	Text      string     `json:"text"`
	URL       string     `json:"url"`
	GifURL    string     `json:"gif_url"`
	Mtime     int        `json:"mtime"`
	Type      int        `json:"type"`
	Meta      EmojiMeta0 `json:"meta"`
	WebpURL   string     `json:"webp_url"`
}

type EmojiPackage struct {
	ID     int           `json:"id"`
	Text   string        `json:"text"`
	URL    string        `json:"url"`
	Mtime  int           `json:"mtime"`
	Type   int           `json:"type"`
	Attr   int           `json:"attr"`
	Meta   EmojiMeta     `json:"meta"`
	Emotes []EmojiEmotes `json:"emotes"`
}

type EmojiData struct {
	Package EmojiPackage `json:"package"`
}

type DownloadInfo struct {
	URL      string
	PkgName  string
	FileName string
}
