## Task1 - API calls
### Programmers Notes
-In this implementation I didn't used any configuration file or models file since the implementation was easy.

-Also I didn't wrote any unit test because it's just a http client and I can reasonably sure that HttpClient works, in a general sense. So there is essentially no real code to test. I think unit tests should be written for deterministic programs that have a complex logic.

-I think the tricky part about this challenge was "The server is able to process requests in parallel." So I called it with the go routine and sent results in a channel.
