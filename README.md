# Fly
Send SMS using [Twilio](https://www.twilio.com/sms/api) via the shell

## Usage

Download it and build it... then...

Set a few environment var's:

```zsh
FLY_TO=2343242424234
FLY_FROM=23432423234
FLY_TWILIO_ACCOUNT=your_twilio_account_number
FLY_TWILIO_AUTH_TOKEN=your_auth_token
```

... or, you can also just pass some arguments to `fly`:

- `-m` - Sets the message
- `-t` - The telephone number you're sending to
- `-f` - From telephone number
- `-a` - Twilio account number
- `-o` - Twilio auth token

Send a message:

```zsh
fly -m "Yes yes yo!"
```


## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## License

The MIT License (MIT)

Copyright (c) 2014 James Duncombe

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
