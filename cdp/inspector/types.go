package inspector

// AUTOGENERATED. DO NOT EDIT.

import (
	"errors"

	. "github.com/knq/chromedp/cdp"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

var (
	_ BackendNode
	_ BackendNodeID
	_ ComputedProperty
	_ ErrorType
	_ Frame
	_ FrameID
	_ LoaderID
	_ Message
	_ MessageError
	_ MethodType
	_ Node
	_ NodeID
	_ NodeType
	_ PseudoType
	_ RGBA
	_ ShadowRootType
	_ Timestamp
)

// DetachReason detach reason.
type DetachReason string

// String returns the DetachReason as string value.
func (t DetachReason) String() string {
	return string(t)
}

// DetachReason values.
const (
	DetachReasonTargetClosed         DetachReason = "target_closed"
	DetachReasonCanceledByUser       DetachReason = "canceled_by_user"
	DetachReasonReplacedWithDevtools DetachReason = "replaced_with_devtools"
	DetachReasonRenderProcessGone    DetachReason = "Render process gone."
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t DetachReason) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t DetachReason) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *DetachReason) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch DetachReason(in.String()) {
	case DetachReasonTargetClosed:
		*t = DetachReasonTargetClosed
	case DetachReasonCanceledByUser:
		*t = DetachReasonCanceledByUser
	case DetachReasonReplacedWithDevtools:
		*t = DetachReasonReplacedWithDevtools
	case DetachReasonRenderProcessGone:
		*t = DetachReasonRenderProcessGone

	default:
		in.AddError(errors.New("unknown DetachReason value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *DetachReason) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}