# Markov chat generator for WhatsApp

[![.github/workflows/release.yaml](https://github.com/2zqa/whatsapp-markov-chat/actions/workflows/release.yaml/badge.svg)](https://github.com/2zqa/whatsapp-markov-chat/actions/workflows/release.yaml)

Generates messages from an exported WhatsApp chat using [Markov Chains](https://en.wikipedia.org/wiki/Markov_chain)

## Description

The program creates a Markov chain for every participant in a chat export, trained on only their messages. Whenever enter is pressed, it prints a message for every author.

## Getting Started

### Dependencies

* Go

### Installing

* `git clone https://github.com/2zqa/whatsapp-markov-chat.git`
* `cd whatsapp-markov-chat`
* `go install`

### Executing program

* [Export your WhatsApp chat](https://faq.whatsapp.com/android/chats/how-to-save-your-chat-history)
* `wamarkov -file <path-to-whatsapp-export.txt>`

## Help

**Q:** No messages are printed!
**A:** Make sure the export is formatted correctly. If it is, please create an issue.

## License

Markov chat generator for WhatsApp is licensed under the [MIT License](LICENSE).
