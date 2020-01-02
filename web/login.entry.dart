import 'dart:html';

import 'package:OAUTH2.APP/loginform.dart';

void main() {
  print("Running Login.Entry");
  window.localStorage['return'] = getParameterByName("return");

  new LoginForm("#frmLogin", "#txtIdentity", "#txtPassword", "#btnLogin");
}

String getParameterByName(String name) {
  final url = window.location.href;
  var uri = Uri.parse(url);
  return uri.queryParameters[name];
}