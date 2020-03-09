class Location {
  final String path;
  final Map<String, String> parameters;

  Location(this.path, this.parameters);

  String parameterAsString() {
    return parameters.keys.map((k) => '$k=${parameters[k]}').join('&');
  }
}