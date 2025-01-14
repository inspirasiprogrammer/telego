package telegohandler

import (
	"regexp"
	"strings"

	"github.com/mymmrac/telego"
)

// Union is true if at least one of predicates is true
func Union(predicates ...Predicate) Predicate {
	return func(update telego.Update) bool {
		for _, p := range predicates {
			if p(update) {
				return true
			}
		}
		return false
	}
}

// Not is true if predicate is false
func Not(predicate Predicate) Predicate {
	return func(update telego.Update) bool {
		return !predicate(update)
	}
}

func anyMassage(message *telego.Message) bool {
	return message != nil
}

// AnyMessage is true if message isn't nil
func AnyMessage() Predicate {
	return func(update telego.Update) bool {
		return anyMassage(update.Message)
	}
}

func anyMassageWithText(message *telego.Message) bool {
	return message != nil && message.Text != ""
}

// AnyMessageWithText is true if message isn't nil and its text not empty
func AnyMessageWithText() Predicate {
	return func(update telego.Update) bool {
		return anyMassageWithText(update.Message)
	}
}

func textEqual(message *telego.Message, text string) bool {
	return message != nil && message.Text == text
}

// TextEqual is true if message isn't nil, and its text equal to specified text
func TextEqual(text string) Predicate {
	return func(update telego.Update) bool {
		return textEqual(update.Message, text)
	}
}

func textEqualFold(message *telego.Message, text string) bool {
	return message != nil && strings.EqualFold(message.Text, text)
}

// TextEqualFold is true if message isn't nil, and its text equal fold (more general form of case-insensitivity equal)
// to specified text
func TextEqualFold(text string) Predicate {
	return func(update telego.Update) bool {
		return textEqualFold(update.Message, text)
	}
}

func textContains(message *telego.Message, text string) bool {
	return message != nil && strings.Contains(message.Text, text)
}

// TextContains is true if message isn't nil, and its text contains specified text
func TextContains(text string) Predicate {
	return func(update telego.Update) bool {
		return textContains(update.Message, text)
	}
}

func textPrefix(message *telego.Message, prefix string) bool {
	return message != nil && strings.HasPrefix(message.Text, prefix)
}

// TextPrefix is true if message isn't nil, and its text has specified prefix
func TextPrefix(prefix string) Predicate {
	return func(update telego.Update) bool {
		return textPrefix(update.Message, prefix)
	}
}

func textSuffix(message *telego.Message, suffix string) bool {
	return message != nil && strings.HasSuffix(message.Text, suffix)
}

// TextSuffix is true if message isn't nil, and its text has specified suffix
func TextSuffix(suffix string) Predicate {
	return func(update telego.Update) bool {
		return textSuffix(update.Message, suffix)
	}
}

func textMatches(message *telego.Message, pattern *regexp.Regexp) bool {
	return message != nil && pattern.MatchString(message.Text)
}

// TextMatches is true if message isn't nil, and its text matches specified regexp
func TextMatches(pattern *regexp.Regexp) Predicate {
	return func(update telego.Update) bool {
		return textMatches(update.Message, pattern)
	}
}

// CommandRegexp matches to command and has match groups on command and arguments
var CommandRegexp = regexp.MustCompile(`^/(\w+) ?(.*)$`)

// CommandMatchGroupsLen represents length of match groups of CommandRegexp
const CommandMatchGroupsLen = 3

// AnyCommand is true if message isn't nil, and it matches to command regexp
func AnyCommand() Predicate {
	return func(update telego.Update) bool {
		return update.Message != nil && CommandRegexp.MatchString(update.Message.Text)
	}
}

// CommandEqual is true if message isn't nil, and it contains specified command
func CommandEqual(command string) Predicate {
	return func(update telego.Update) bool {
		if update.Message == nil {
			return false
		}

		matches := CommandRegexp.FindStringSubmatch(update.Message.Text)
		if len(matches) != CommandMatchGroupsLen {
			return false
		}

		return strings.EqualFold(matches[1], command)
	}
}

// CommandEqualArgc is true if message isn't nil, and it contains specified command with number of args
func CommandEqualArgc(command string, argc int) Predicate {
	return func(update telego.Update) bool {
		if update.Message == nil {
			return false
		}

		matches := CommandRegexp.FindStringSubmatch(update.Message.Text)
		if len(matches) != CommandMatchGroupsLen {
			return false
		}

		return strings.EqualFold(matches[1], command) &&
			(argc == 0 && matches[2] == "" || len(strings.Split(matches[2], " ")) == argc)
	}
}

// CommandEqualArgv is true if message isn't nil, and it contains specified command and args
func CommandEqualArgv(command string, argv ...string) Predicate {
	return func(update telego.Update) bool {
		if update.Message == nil {
			return false
		}

		matches := CommandRegexp.FindStringSubmatch(update.Message.Text)
		if len(matches) != CommandMatchGroupsLen {
			return false
		}

		return strings.EqualFold(matches[1], command) &&
			(len(argv) == 0 && matches[2] == "" || matches[2] == strings.Join(argv, " "))
	}
}

// AnyEditedMessage is true if edited message isn't nil
func AnyEditedMessage() Predicate {
	return func(update telego.Update) bool {
		return anyMassage(update.EditedMessage)
	}
}

// AnyEditedMessageWithText is true if edited message isn't nil and its text not empty
func AnyEditedMessageWithText() Predicate {
	return func(update telego.Update) bool {
		return anyMassageWithText(update.EditedMessage)
	}
}

// EditedTextEqual is true if edited message isn't nil, and its text equal to specified text
func EditedTextEqual(text string) Predicate {
	return func(update telego.Update) bool {
		return textEqual(update.EditedMessage, text)
	}
}

// EditedTextEqualFold is true if edited message isn't nil, and its text equal fold (more general form of
// case-insensitivity equal) to specified text
func EditedTextEqualFold(text string) Predicate {
	return func(update telego.Update) bool {
		return textEqualFold(update.EditedMessage, text)
	}
}

// EditedTextContains is true if edited message isn't nil, and its text contains specified text
func EditedTextContains(text string) Predicate {
	return func(update telego.Update) bool {
		return textContains(update.EditedMessage, text)
	}
}

// EditedTextPrefix is true if edited message isn't nil, and its text has specified prefix
func EditedTextPrefix(prefix string) Predicate {
	return func(update telego.Update) bool {
		return textPrefix(update.EditedMessage, prefix)
	}
}

// EditedTextSuffix is true if edited message isn't nil, and its text has specified suffix
func EditedTextSuffix(suffix string) Predicate {
	return func(update telego.Update) bool {
		return textSuffix(update.EditedMessage, suffix)
	}
}

// EditedTextMatches is true if edited message isn't nil, and its text matches specified regexp
func EditedTextMatches(pattern *regexp.Regexp) Predicate {
	return func(update telego.Update) bool {
		return textMatches(update.EditedMessage, pattern)
	}
}

// AnyChannelPost is true if channel post isn't nil
func AnyChannelPost() Predicate {
	return func(update telego.Update) bool {
		return anyMassage(update.ChannelPost)
	}
}

// AnyChannelPostWithText is true if channel post isn't nil and its text not empty
func AnyChannelPostWithText() Predicate {
	return func(update telego.Update) bool {
		return anyMassageWithText(update.ChannelPost)
	}
}

// PostTextEqual is true if channel post isn't nil, and its text equal to specified text
func PostTextEqual(text string) Predicate {
	return func(update telego.Update) bool {
		return textEqual(update.ChannelPost, text)
	}
}

// PostTextEqualFold is true if channel post isn't nil, and its text equal fold (more general form of case-insensitivity
// equal) to specified text
func PostTextEqualFold(text string) Predicate {
	return func(update telego.Update) bool {
		return textEqualFold(update.ChannelPost, text)
	}
}

// PostTextContains is true if channel post isn't nil, and its text contains specified text
func PostTextContains(text string) Predicate {
	return func(update telego.Update) bool {
		return textContains(update.ChannelPost, text)
	}
}

// PostTextPrefix is true if channel post isn't nil, and its text has specified prefix
func PostTextPrefix(prefix string) Predicate {
	return func(update telego.Update) bool {
		return textPrefix(update.ChannelPost, prefix)
	}
}

// PostTextSuffix is true if channel post isn't nil, and its text has specified suffix
func PostTextSuffix(suffix string) Predicate {
	return func(update telego.Update) bool {
		return textSuffix(update.ChannelPost, suffix)
	}
}

// PostTextMatches is true if channel post isn't nil, and its text matches specified regexp
func PostTextMatches(pattern *regexp.Regexp) Predicate {
	return func(update telego.Update) bool {
		return textMatches(update.ChannelPost, pattern)
	}
}

// AnyEditedChannelPost is true if edited channel post isn't nil
func AnyEditedChannelPost() Predicate {
	return func(update telego.Update) bool {
		return anyMassage(update.EditedChannelPost)
	}
}

// AnyEditedChannelPostWithText is true if edited channel post isn't nil and its text not empty
func AnyEditedChannelPostWithText() Predicate {
	return func(update telego.Update) bool {
		return anyMassageWithText(update.EditedChannelPost)
	}
}

// EditedPostTextEqual is true if edited channel post isn't nil, and its text equal to specified text
func EditedPostTextEqual(text string) Predicate {
	return func(update telego.Update) bool {
		return textEqual(update.EditedChannelPost, text)
	}
}

// EditedPostTextEqualFold is true if edited channel post isn't nil, and its text equal fold (more general form of
// case-insensitivity equal) to specified text
func EditedPostTextEqualFold(text string) Predicate {
	return func(update telego.Update) bool {
		return textEqualFold(update.EditedChannelPost, text)
	}
}

// EditedPostTextContains is true if edited channel post isn't nil, and its text contains specified text
func EditedPostTextContains(text string) Predicate {
	return func(update telego.Update) bool {
		return textContains(update.EditedChannelPost, text)
	}
}

// EditedPostTextPrefix is true if edited channel post isn't nil, and its text has specified prefix
func EditedPostTextPrefix(prefix string) Predicate {
	return func(update telego.Update) bool {
		return textPrefix(update.EditedChannelPost, prefix)
	}
}

// EditedPostTextSuffix is true if edited channel post isn't nil, and its text has specified suffix
func EditedPostTextSuffix(suffix string) Predicate {
	return func(update telego.Update) bool {
		return textSuffix(update.EditedChannelPost, suffix)
	}
}

// EditedPostTextMatches is true if edited channel post isn't nil, and its text matches specified regexp
func EditedPostTextMatches(pattern *regexp.Regexp) Predicate {
	return func(update telego.Update) bool {
		return textMatches(update.EditedChannelPost, pattern)
	}
}

// AnyInlineQuery is true if inline query isn't nil
func AnyInlineQuery() Predicate {
	return func(update telego.Update) bool {
		return update.InlineQuery != nil
	}
}

// InlineQueryEqual is true if inline query isn't nil, and its query equal to specified text
func InlineQueryEqual(text string) Predicate {
	return func(update telego.Update) bool {
		return update.InlineQuery != nil && update.InlineQuery.Query == text
	}
}

// InlineQueryEqualFold is true if inline query isn't nil, and its query equal fold (more general form of
// case-insensitivity equal) to specified text
func InlineQueryEqualFold(text string) Predicate {
	return func(update telego.Update) bool {
		return update.InlineQuery != nil && strings.EqualFold(update.InlineQuery.Query, text)
	}
}

// InlineQueryContains is true if inline query isn't nil, and its query contains specified text
func InlineQueryContains(text string) Predicate {
	return func(update telego.Update) bool {
		return update.InlineQuery != nil && strings.Contains(update.InlineQuery.Query, text)
	}
}

// InlineQueryPrefix is true if inline query isn't nil, and its query has specified prefix
func InlineQueryPrefix(prefix string) Predicate {
	return func(update telego.Update) bool {
		return update.InlineQuery != nil && strings.HasPrefix(update.InlineQuery.Query, prefix)
	}
}

// InlineQuerySuffix is true if inline query isn't nil, and its query has specified suffix
func InlineQuerySuffix(suffix string) Predicate {
	return func(update telego.Update) bool {
		return update.InlineQuery != nil && strings.HasSuffix(update.InlineQuery.Query, suffix)
	}
}

// InlineQueryMatches is true if inline query isn't nil, and its query matches specified regexp
func InlineQueryMatches(pattern *regexp.Regexp) Predicate {
	return func(update telego.Update) bool {
		return update.InlineQuery != nil && pattern.MatchString(update.InlineQuery.Query)
	}
}

// AnyChosenInlineResult is true if chosen inline result isn't nil
func AnyChosenInlineResult() Predicate {
	return func(update telego.Update) bool {
		return update.ChosenInlineResult != nil
	}
}

// AnyCallbackQuery is true if callback query isn't nil
func AnyCallbackQuery() Predicate {
	return func(update telego.Update) bool {
		return update.CallbackQuery != nil
	}
}

// AnyCallbackQueryWithMessage is true if callback query and its message isn't nil
func AnyCallbackQueryWithMessage() Predicate {
	return func(update telego.Update) bool {
		return update.CallbackQuery != nil && update.CallbackQuery.Message != nil
	}
}

// CallbackDataEqual is true if callback query isn't nil, and its data equal to specified text
func CallbackDataEqual(text string) Predicate {
	return func(update telego.Update) bool {
		return update.CallbackQuery != nil && update.CallbackQuery.Data == text
	}
}

// CallbackDataEqualFold is true if callback query isn't nil, and its data equal fold (more general form of
// case-insensitivity equal) to specified text
func CallbackDataEqualFold(text string) Predicate {
	return func(update telego.Update) bool {
		return update.CallbackQuery != nil && strings.EqualFold(update.CallbackQuery.Data, text)
	}
}

// CallbackDataContains is true if callback query isn't nil, and its data contains specified text
func CallbackDataContains(text string) Predicate {
	return func(update telego.Update) bool {
		return update.CallbackQuery != nil && strings.Contains(update.CallbackQuery.Data, text)
	}
}

// CallbackDataPrefix is true if callback query isn't nil, and its data has specified prefix
func CallbackDataPrefix(prefix string) Predicate {
	return func(update telego.Update) bool {
		return update.CallbackQuery != nil && strings.HasPrefix(update.CallbackQuery.Data, prefix)
	}
}

// CallbackDataSuffix is true if callback query isn't nil, and its data has specified suffix
func CallbackDataSuffix(suffix string) Predicate {
	return func(update telego.Update) bool {
		return update.CallbackQuery != nil && strings.HasSuffix(update.CallbackQuery.Data, suffix)
	}
}

// CallbackDataMatches is true if callback query isn't nil, and its data matches specified regexp
func CallbackDataMatches(pattern *regexp.Regexp) Predicate {
	return func(update telego.Update) bool {
		return update.CallbackQuery != nil && pattern.MatchString(update.CallbackQuery.Data)
	}
}

// AnyShippingQuery is true if shipping query isn't nil
func AnyShippingQuery() Predicate {
	return func(update telego.Update) bool {
		return update.ShippingQuery != nil
	}
}

// AnyPreCheckoutQuery is true if pre checkout query isn't nil
func AnyPreCheckoutQuery() Predicate {
	return func(update telego.Update) bool {
		return update.PreCheckoutQuery != nil
	}
}

// AnyPoll is true if poll isn't nil
func AnyPoll() Predicate {
	return func(update telego.Update) bool {
		return update.Poll != nil
	}
}

// AnyPollAnswer is true if poll answer isn't nil
func AnyPollAnswer() Predicate {
	return func(update telego.Update) bool {
		return update.PollAnswer != nil
	}
}

// AnyMyChatMember is true if my chat member isn't nil
func AnyMyChatMember() Predicate {
	return func(update telego.Update) bool {
		return update.MyChatMember != nil
	}
}

// AnyChatMember is true if chat member isn't nil
func AnyChatMember() Predicate {
	return func(update telego.Update) bool {
		return update.ChatMember != nil
	}
}

// AnyChatJoinRequest is true if chat join request isn't nil
func AnyChatJoinRequest() Predicate {
	return func(update telego.Update) bool {
		return update.ChatJoinRequest != nil
	}
}
