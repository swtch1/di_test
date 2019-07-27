package di_test
//
//import (
//	"github.com/rs/zerolog"
//	"github.com/rs/zerolog/log"
//)

//type Log struct{}

//// Print is a mirror of zerolog's log.Print().
//func (_ *Log) Print(v ...interface{}) {
//	log.Print(v)
//}
//
//// Printf is a mirror of zerolog's log.Printf().
//// Printf sends a log event using debug level and no extra field.
//// Arguments are handled in the manner of fmt.Printf.
//func (_ *Log) Printf(format string, v ...interface{}) {
//	log.Printf(format, v)
//}
//
//// Debug is a mirror of zerolog's log.Debug().
//// Debug starts a new message with debug level.
////
//// You must call Msg on the returned event in order to send the event.
//func (_ *Log) Debug() *zerolog.Event {
//	return log.Debug()
//}
//
//// Info is a mirror of zerolog's log.Info().
//// Info starts a new message with info level.
////
//// You must call Msg on the returned event in order to send the event.
//func (_ *Log) Info() *zerolog.Event {
//	return log.Info()
//}
//
//// Warn is a mirror of zerolog's log.Warn().
//// Warn starts a new message with warn level.
////
//// You must call Msg on the returned event in order to send the event.
//func (_ *Log) Warn() *zerolog.Event {
//	return log.Warn()
//}
//
//// Error is a mirror of zerolog's log.Error().
//// Error starts a new message with error level.
////
//// You must call Msg on the returned event in order to send the event.
//func (_ *Log) Error() *zerolog.Event {
//	return log.Error()
//}
//
//// Fatal is a mirror of zerolog's log.Fatal().
//// Fatal starts a new message with fatal level. The os.Exit(1) function
//// is called by the Msg method.
////
//// You must call Msg on the returned event in order to send the event.
//func (_ *Log) Fatal() *zerolog.Event {
//	return log.Fatal()
//}
//
//// Panic is a mirror of zerolog's log.Panic().
//// Panic starts a new message with panic level. The message is also sent
//// to the panic function.
////
//// You must call Msg on the returned event in order to send the event.
//func (_ *Log) Panic() *zerolog.Event {
//	return log.Panic()
//}
