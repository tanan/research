import 'dart:async';

class HttpErrorHandler {
  final StreamController<HttpError> _controller = StreamController<HttpError>();
  Stream<HttpError> _stream;
  final Map<OnHttpError, StreamSubscription<HttpError>> _subscriptions = {};

  HttpErrorHandler() {
    _stream = _controller.stream;
  }

  void subscribe(OnHttpError handler) {
    _subscriptions[handler] = _stream.listen(handler);
  }

  void unsubscribe(OnHttpError handler) {
    _subscriptions[handler]?.cancel();
  }

  void notify(HttpError error) {
    _controller.add(error);
  }
}

typedef OnHttpError = void Function(HttpError e);

class HttpError {
  final int statusCode;
  final String message;

  HttpError(this.statusCode, { this.message} );


  @override
  String toString() {
    return 'HttpError(${statusCode}, ${message})';
  }
}