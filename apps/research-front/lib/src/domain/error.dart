class Error {
  final String message;

  Error(this.message);
}

typedef ErrorHandler = void Function(Error e);