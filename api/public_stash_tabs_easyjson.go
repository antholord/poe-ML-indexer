// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package api

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson8cf7917eDecodeGithubComAntholordPoeMLIndexerApi(in *jlexer.Lexer, out *PublicStashTabs) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "next_change_id":
			out.NextChangeId = string(in.String())
		case "stashes":
			if in.IsNull() {
				in.Skip()
				out.Stashes = nil
			} else {
				in.Delim('[')
				if out.Stashes == nil {
					if !in.IsDelim(']') {
						out.Stashes = make([]Stash, 0, 1)
					} else {
						out.Stashes = []Stash{}
					}
				} else {
					out.Stashes = (out.Stashes)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Stash
					easyjson8cf7917eDecodeGithubComAntholordPoeMLIndexerApi1(in, &v1)
					out.Stashes = append(out.Stashes, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson8cf7917eEncodeGithubComAntholordPoeMLIndexerApi(out *jwriter.Writer, in PublicStashTabs) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"next_change_id\":")
	out.String(string(in.NextChangeId))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"stashes\":")
	if in.Stashes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in.Stashes {
			if v2 > 0 {
				out.RawByte(',')
			}
			easyjson8cf7917eEncodeGithubComAntholordPoeMLIndexerApi1(out, v3)
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PublicStashTabs) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8cf7917eEncodeGithubComAntholordPoeMLIndexerApi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PublicStashTabs) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8cf7917eEncodeGithubComAntholordPoeMLIndexerApi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PublicStashTabs) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8cf7917eDecodeGithubComAntholordPoeMLIndexerApi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PublicStashTabs) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8cf7917eDecodeGithubComAntholordPoeMLIndexerApi(l, v)
}
func easyjson8cf7917eDecodeGithubComAntholordPoeMLIndexerApi1(in *jlexer.Lexer, out *Stash) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "accountName":
			out.AccountName = string(in.String())
		case "lastCharacterName":
			out.LastCharacterName = string(in.String())
		case "id":
			out.Id = string(in.String())
		case "stash":
			out.Label = string(in.String())
		case "stashType":
			out.Type = string(in.String())
		case "items":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]Item, 0, 1)
					} else {
						out.Items = []Item{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v4 Item
					easyjson8cf7917eDecodeGithubComAntholordPoeMLIndexerApi2(in, &v4)
					out.Items = append(out.Items, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "public":
			out.IsPublic = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson8cf7917eEncodeGithubComAntholordPoeMLIndexerApi1(out *jwriter.Writer, in Stash) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"accountName\":")
	out.String(string(in.AccountName))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"lastCharacterName\":")
	out.String(string(in.LastCharacterName))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"id\":")
	out.String(string(in.Id))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"stash\":")
	out.String(string(in.Label))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"stashType\":")
	out.String(string(in.Type))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"items\":")
	if in.Items == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in.Items {
			if v5 > 0 {
				out.RawByte(',')
			}
			easyjson8cf7917eEncodeGithubComAntholordPoeMLIndexerApi2(out, v6)
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"public\":")
	out.Bool(bool(in.IsPublic))
	out.RawByte('}')
}
func easyjson8cf7917eDecodeGithubComAntholordPoeMLIndexerApi2(in *jlexer.Lexer, out *Item) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "typeLine":
			out.Type = string(in.String())
		case "explicitMods":
			if in.IsNull() {
				in.Skip()
				out.ExplicitMods = nil
			} else {
				in.Delim('[')
				if out.ExplicitMods == nil {
					if !in.IsDelim(']') {
						out.ExplicitMods = make([]string, 0, 4)
					} else {
						out.ExplicitMods = []string{}
					}
				} else {
					out.ExplicitMods = (out.ExplicitMods)[:0]
				}
				for !in.IsDelim(']') {
					var v7 string
					v7 = string(in.String())
					out.ExplicitMods = append(out.ExplicitMods, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "implicitMods":
			if in.IsNull() {
				in.Skip()
				out.ImplicitMods = nil
			} else {
				in.Delim('[')
				if out.ImplicitMods == nil {
					if !in.IsDelim(']') {
						out.ImplicitMods = make([]string, 0, 4)
					} else {
						out.ImplicitMods = []string{}
					}
				} else {
					out.ImplicitMods = (out.ImplicitMods)[:0]
				}
				for !in.IsDelim(']') {
					var v8 string
					v8 = string(in.String())
					out.ImplicitMods = append(out.ImplicitMods, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "utilityMods":
			if in.IsNull() {
				in.Skip()
				out.UtilityMods = nil
			} else {
				in.Delim('[')
				if out.UtilityMods == nil {
					if !in.IsDelim(']') {
						out.UtilityMods = make([]string, 0, 4)
					} else {
						out.UtilityMods = []string{}
					}
				} else {
					out.UtilityMods = (out.UtilityMods)[:0]
				}
				for !in.IsDelim(']') {
					var v9 string
					v9 = string(in.String())
					out.UtilityMods = append(out.UtilityMods, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "enchantMods":
			if in.IsNull() {
				in.Skip()
				out.EnchantMods = nil
			} else {
				in.Delim('[')
				if out.EnchantMods == nil {
					if !in.IsDelim(']') {
						out.EnchantMods = make([]string, 0, 4)
					} else {
						out.EnchantMods = []string{}
					}
				} else {
					out.EnchantMods = (out.EnchantMods)[:0]
				}
				for !in.IsDelim(']') {
					var v10 string
					v10 = string(in.String())
					out.EnchantMods = append(out.EnchantMods, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "craftedMods":
			if in.IsNull() {
				in.Skip()
				out.CraftedMods = nil
			} else {
				in.Delim('[')
				if out.CraftedMods == nil {
					if !in.IsDelim(']') {
						out.CraftedMods = make([]string, 0, 4)
					} else {
						out.CraftedMods = []string{}
					}
				} else {
					out.CraftedMods = (out.CraftedMods)[:0]
				}
				for !in.IsDelim(']') {
					var v11 string
					v11 = string(in.String())
					out.CraftedMods = append(out.CraftedMods, v11)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "cosmeticMods":
			if in.IsNull() {
				in.Skip()
				out.CosmeticMods = nil
			} else {
				in.Delim('[')
				if out.CosmeticMods == nil {
					if !in.IsDelim(']') {
						out.CosmeticMods = make([]string, 0, 4)
					} else {
						out.CosmeticMods = []string{}
					}
				} else {
					out.CosmeticMods = (out.CosmeticMods)[:0]
				}
				for !in.IsDelim(']') {
					var v12 string
					v12 = string(in.String())
					out.CosmeticMods = append(out.CosmeticMods, v12)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "note":
			out.Note = string(in.String())
		case "verified":
			out.IsVerified = bool(in.Bool())
		case "ilvl":
			out.ItemLevel = int(in.Int())
		case "league":
			out.League = string(in.String())
		case "id":
			out.Id = string(in.String())
		case "identified":
			out.IsIdentified = bool(in.Bool())
		case "frameType":
			out.FrameType = FrameType(in.Int())
		case "stackSize":
			out.StackSize = int(in.Int())
		case "maxStackSize":
			out.MaxStackSize = int(in.Int())
		case "inventoryId":
			out.InventoryId = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson8cf7917eEncodeGithubComAntholordPoeMLIndexerApi2(out *jwriter.Writer, in Item) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"typeLine\":")
	out.String(string(in.Type))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"explicitMods\":")
	if in.ExplicitMods == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v13, v14 := range in.ExplicitMods {
			if v13 > 0 {
				out.RawByte(',')
			}
			out.String(string(v14))
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"implicitMods\":")
	if in.ImplicitMods == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v15, v16 := range in.ImplicitMods {
			if v15 > 0 {
				out.RawByte(',')
			}
			out.String(string(v16))
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"utilityMods\":")
	if in.UtilityMods == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v17, v18 := range in.UtilityMods {
			if v17 > 0 {
				out.RawByte(',')
			}
			out.String(string(v18))
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"enchantMods\":")
	if in.EnchantMods == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v19, v20 := range in.EnchantMods {
			if v19 > 0 {
				out.RawByte(',')
			}
			out.String(string(v20))
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"craftedMods\":")
	if in.CraftedMods == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v21, v22 := range in.CraftedMods {
			if v21 > 0 {
				out.RawByte(',')
			}
			out.String(string(v22))
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"cosmeticMods\":")
	if in.CosmeticMods == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v23, v24 := range in.CosmeticMods {
			if v23 > 0 {
				out.RawByte(',')
			}
			out.String(string(v24))
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"note\":")
	out.String(string(in.Note))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"verified\":")
	out.Bool(bool(in.IsVerified))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"ilvl\":")
	out.Int(int(in.ItemLevel))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"league\":")
	out.String(string(in.League))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"id\":")
	out.String(string(in.Id))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"identified\":")
	out.Bool(bool(in.IsIdentified))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"frameType\":")
	out.Int(int(in.FrameType))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"stackSize\":")
	out.Int(int(in.StackSize))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"maxStackSize\":")
	out.Int(int(in.MaxStackSize))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"inventoryId\":")
	out.String(string(in.InventoryId))
	out.RawByte('}')
}
