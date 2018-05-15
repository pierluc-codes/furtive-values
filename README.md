# Furtive value

Furtive value is software that transform encryped value into their appropriate decrypted value.
The magic is handled with a embedded data structure into the software.

## Dependencies

* vgo
* Google Cloud Project

## How it works

Furtivate value is processing each line one by one.
When it finds a specific pattern of string, it decoded the embeded payload.
The embedded payload is then used to execute a decryption call with Google Cloud KMS.

### Secret string

The secret string is expected to be the following format

```
(furtivevalue1:my-embeded-payload)
```

Example:

```(furtivevalue1:eyJiIjoiZ2NwIiwicCI6InNvbGlkLW11c2UtMjAzOTAxIiwibCI6Imdsb2JhbCIsInIiOiJzdGFnaW5nLWtleXJpbmciLCJrIjoia2V5LW51bWJlci1vbmUiLCJjIjoiQ2lRQWtheDU0UGxXQS8zMG5LSnVHaXIvODVTS2lsaS9nZUlzVUJzNkpMcmhEa1lzMERZU1R3RC9qcjFtS2FNdG1HdFhpTVMzVzd2N3ZrUVBHYTVlbitLY3pQSGt5a2hocG9aRHFWYm4vWUNxalc2QWRsMThvbFpyVm9hRmtubWUzVzYzYWlWN2tYVCtveE5uYXpwZ0t4dmlhWGpMS0cwPSJ9)
```

The value of my-embeded-payload is base64 encoded json.

### Embeded payload

The base64 encoded json is the following format

```
{
    "b": "" 
	"p": ""
	"l": ""
	"r": ""
	"k": ""
	"c": ""
}
```

field | definition
 ------------------|------------------
b | The backend the use to decrypt. At the moment, only the value 'gcp' is supported.
p | Name of the project on Google Cloud platform
l | The location of the Google Cloud KMS keyring
r | The keyring of the Google Cloud KMS that contains the key to decrypt
k | The key f the Google Cloud KMS to use to decrypt the cipher text
c | The cipher text to decrypt

Example:
{
   "b":"gcp",
   "p":"solid-muse-203901",
   "l":"global",
   "r":"staging-keyring",
   "k":"key-number-one",
   "c":"CiQAkax54PlWA/30nKJuGir/85SKili/geIsUBs6JLrhDkYs0DYSTwD/jr1mKaMtmGtXiMS3W7v7vkQPGa5en+KczPHkykhhpoZDqVbn/YCqjW6Adl18olZrVoaFknme3W63aiV7kXT+oxNnazpgKxviaXjLKG0="
}

## Development

### Quick running 

```
cat textfile.txt | vgo run main.go
```

### Testing

```
vgo test ./...
```

### Building

```
vgo build
```