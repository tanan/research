abstract class FCC<T> extends Iterable<T> {
  List<T> get values;

  @override
  Iterator<T> get iterator => values.iterator;

  List<R> mapToList<R>(R Function(T each) f) => map(f).toList();
}