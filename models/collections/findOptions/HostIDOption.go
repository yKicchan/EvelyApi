package findOptions

type HostIDOption interface {
	FindOption
	SetHostID(string)
	GetHostID() string
	IsHostIDSet() bool
}

type hostIDOption struct {
	// イベントの作成者ID
	hostID string
}

/**
 * 設定されたイベントの作成者IDを返す
 * @return string イベントの作成者ID
 */
func (this *hostIDOption) GetHostID() string { return this.hostID }

/**
 * 検索オプションにイベントの作成者IDを設定する
 * @param id 作成者ID
 */
func (this *hostIDOption) SetHostID(id string) { this.hostID = id }

/**
 * 作成者IDでのイベント検索が有効化を調べる
 * @return bool
 */
func (this *hostIDOption) IsHostIDSet() bool { return this.hostID != "" }
