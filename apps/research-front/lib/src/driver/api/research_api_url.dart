
const String BASE_URL = String.fromEnvironment('RESEARCH_API_BASE', defaultValue: 'http://localhost:8000');

class ResearchApiURL {
  String latest() => '$BASE_URL/v1/articles';

  String article(String id) => '$BASE_URL/v1/articles/$id';
}