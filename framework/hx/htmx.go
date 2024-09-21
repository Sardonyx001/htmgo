package hx

type Attribute = string
type Header = string
type Event = string

// https://htmx.org/reference/#events
const (
	GetAttr         Attribute = "hx-get"
	PostAttr        Attribute = "hx-post"
	PushUrlAttr     Attribute = "hx-push-url"
	SelectAttr      Attribute = "hx-select"
	SelectOobAttr   Attribute = "hx-select-oob"
	SwapAttr        Attribute = "hx-swap"
	SwapOobAttr     Attribute = "hx-swap-oob"
	TargetAttr      Attribute = "hx-target"
	TriggerAttr     Attribute = "hx-trigger"
	ValsAttr        Attribute = "hx-vals"
	BoostAttr       Attribute = "hx-boost"
	ConfirmAttr     Attribute = "hx-confirm"
	DeleteAttr      Attribute = "hx-delete"
	DisableAttr     Attribute = "hx-disable"
	DisabledEltAttr Attribute = "hx-disabled-elt"
	DisinheritAttr  Attribute = "hx-disinherit"
	EncodingAttr    Attribute = "hx-encoding"
	ExtAttr         Attribute = "hx-ext"
	HeadersAttr     Attribute = "hx-headers"
	HistoryAttr     Attribute = "hx-history"
	HistoryEltAttr  Attribute = "hx-history-elt"
	IncludeAttr     Attribute = "hx-include"
	IndicatorAttr   Attribute = "hx-indicator"
	InheritAttr     Attribute = "hx-inherit"
	ParamsAttr      Attribute = "hx-params"
	PatchAttr       Attribute = "hx-patch"
	PreserveAttr    Attribute = "hx-preserve"
	PromptAttr      Attribute = "hx-prompt"
	PutAttr         Attribute = "hx-put"
	ReplaceUrlAttr  Attribute = "hx-replace-url"
	RequestAttr     Attribute = "hx-request"
	SyncAttr        Attribute = "hx-sync"
	ValidateAttr    Attribute = "hx-validate"
)

const (
	LocationHeader           Header = "HX-Location"
	PushUrlHeader            Header = "HX-Push-Url"
	RedirectHeader           Header = "HX-Redirect"
	RefreshHeader            Header = "HX-Refresh"
	ReplaceUrlHeader         Header = "HX-Replace-Url"
	CurrentUrlHeader         Header = "HX-Current-Url"
	ReswapHeader             Header = "HX-Reswap"
	RetargetHeader           Header = "HX-Retarget"
	ReselectHeader           Header = "HX-Reselect"
	TriggerHeader            Header = "HX-Trigger"
	TriggerAfterSettleHeader Header = "HX-Trigger-After-Settle"
	TriggerAfterSwapHeader   Header = "HX-Trigger-After-Swap"
)

const (
	// AbortEvent Htmx Events
	AbortEvent                 Event = "htmx:abort"
	AfterOnLoadEvent           Event = "htmx:afterOnLoad"
	AfterProcessNodeEvent      Event = "htmx:afterProcessNode"
	AfterRequestEvent          Event = "htmx:afterRequest"
	OnMutationErrorEvent       Event = "htmx:onMutationError"
	AfterSettleEvent           Event = "htmx:afterSettle"
	AfterSwapEvent             Event = "htmx:afterSwap"
	BeforeCleanupElementEvent  Event = "htmx:beforeCleanupElement"
	BeforeOnLoadEvent          Event = "htmx:beforeOnLoad"
	BeforeProcessNodeEvent     Event = "htmx:beforeProcessNode"
	BeforeRequestEvent         Event = "htmx:beforeRequest"
	BeforeSwapEvent            Event = "htmx:beforeSwap"
	BeforeSendEvent            Event = "htmx:beforeSend"
	ConfigRequestEvent         Event = "htmx:configRequest"
	ConfirmEvent               Event = "htmx:confirm"
	HistoryCacheErrorEvent     Event = "htmx:historyCacheError"
	HistoryCacheMissEvent      Event = "htmx:historyCacheMiss"
	HistoryCacheMissErrorEvent Event = "htmx:historyCacheMissError"
	HistoryCacheMissLoadEvent  Event = "htmx:historyCacheMissLoad"
	HistoryRestoreEvent        Event = "htmx:historyRestore"
	BeforeHistorySaveEvent     Event = "htmx:beforeHistorySave"
	LoadEvent                  Event = "htmx:load"
	NoSSESourceErrorEvent      Event = "htmx:noSSESourceError"
	OnLoadErrorEvent           Event = "htmx:onLoadError"
	OobAfterSwapEvent          Event = "htmx:oobAfterSwap"
	OobBeforeSwapEvent         Event = "htmx:oobBeforeSwap"
	OobErrorNoTargetEvent      Event = "htmx:oobErrorNoTarget"
	PromptEvent                Event = "htmx:prompt"
	PushedIntoHistoryEvent     Event = "htmx:pushedIntoHistory"
	ResponseErrorEvent         Event = "htmx:responseError"
	SendErrorEvent             Event = "htmx:sendError"
	SSEErrorEvent              Event = "htmx:sseError"
	SSEOpenEvent               Event = "htmx:sseOpen"
	SwapErrorEvent             Event = "htmx:swapError"
	TargetErrorEvent           Event = "htmx:targetError"
	TimeoutEvent               Event = "htmx:timeout"
	ValidationValidateEvent    Event = "htmx:validation:validate"
	ValidationFailedEvent      Event = "htmx:validation:failed"
	ValidationHaltedEvent      Event = "htmx:validation:halted"
	XhrAbortEvent              Event = "htmx:xhr:abort"
	XhrLoadEndEvent            Event = "htmx:xhr:loadend"
	XhrLoadStartEvent          Event = "htmx:xhr:loadstart"
	XhrProgressEvent           Event = "htmx:xhr:progress"

	// RevealedEvent Misc Events
	RevealedEvent   Event = "revealed"
	InstersectEvent Event = "intersect"
	PollingEvent    Event = "every"

	// ClickEvent Dom Events
	ClickEvent       Event = "onclick"
	ChangeEvent      Event = "onchange"
	InputEvent       Event = "oninput"
	FocusEvent       Event = "onfocus"
	BlurEvent        Event = "onblur"
	KeyDownEvent     Event = "onkeydown"
	KeyUpEvent       Event = "onkeyup"
	KeyPressEvent    Event = "onkeypress"
	SubmitEvent      Event = "onsubmit"
	LoadDomEvent     Event = "onload"
	UnloadEvent      Event = "onunload"
	ResizeEvent      Event = "onresize"
	ScrollEvent      Event = "onscroll"
	DblClickEvent    Event = "ondblclick"
	MouseOverEvent   Event = "onmouseover"
	MouseOutEvent    Event = "onmouseout"
	MouseMoveEvent   Event = "onmousemove"
	MouseDownEvent   Event = "onmousedown"
	MouseUpEvent     Event = "onmouseup"
	ContextMenuEvent Event = "oncontextmenu"
	DragStartEvent   Event = "ondragstart"
	DragEvent        Event = "ondrag"
	DragEnterEvent   Event = "ondragenter"
	DragLeaveEvent   Event = "ondragleave"
	DragOverEvent    Event = "ondragover"
	DropEvent        Event = "ondrop"
	DragEndEvent     Event = "ondragend"
)