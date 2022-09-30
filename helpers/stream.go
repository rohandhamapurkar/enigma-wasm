package helpers

import (
	"io"
	"syscall/js"
)

func GetJSReadableStreamFromIOReader(reader io.Reader) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		underlyingSource := map[string]interface{}{
			// start method
			"start": js.FuncOf(func(this js.Value, args []js.Value) interface{} {
				controller := args[0]
				go func() {
					// Read the entire stream and pass it to JavaScript
					for {
						// Read up to 16KB at a time
						buf := make([]byte, 16384)
						n, err := reader.Read(buf)
						if err != nil && err != io.EOF {
							// Tell the controller we have an error
							// We're ignoring "EOF" however, which means the stream was done
							errorConstructor := js.Global().Get("Error")
							errorObject := errorConstructor.New(err.Error())
							controller.Call("error", errorObject)
							return
						}
						if n > 0 {
							// If we read anything, send it to JavaScript using the "enqueue" method on the controller
							// We need to convert it to a Uint8Array first
							arrayConstructor := js.Global().Get("Uint8Array")
							dataJS := arrayConstructor.New(n)
							js.CopyBytesToJS(dataJS, buf[0:n])
							controller.Call("enqueue", dataJS)
						}
						if err == io.EOF {
							// Stream is done, so call the "close" method on the controller
							controller.Call("close")
							return
						}
					}
				}()

				return nil
			}),
			// cancel method
			"cancel": js.FuncOf(func(this js.Value, args []js.Value) interface{} {
				return nil
			}),
		}

		// Create a ReadableStream object from the underlyingSource object
		readableStreamConstructor := js.Global().Get("ReadableStream")
		readableStream := readableStreamConstructor.New(underlyingSource)

		return readableStream
	})
}