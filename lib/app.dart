import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/bodies/app.dart';

Future<App> getApp() async {
  await identifyLocation();

  var appUrl = window.localStorage['return'];
  var ip = await getIP();
  var location = window.localStorage['location'];
  HiddenInputElement instanceElem = querySelector("#InstanceID");

  return new App(appUrl, ip, location, instanceElem.value);
}

//geolocation only works on HTTPS
void identifyLocation() {
  if (window.navigator != null) {
    final geo = window.navigator.geolocation;
    geo.getCurrentPosition(maximumAge: new Duration(hours: 8), timeout: new Duration(seconds: 10)).then(storeLocation, onError: locationFailed);
  }
}

Future<String> getIP() async {
  var resp = await HttpRequest.getString('https://jsonip.com');

  return jsonDecode(resp)["ip"];
}

void storeLocation(position) {
 window.localStorage['location'] =
          '${position.coords.latitude}, ${position.coords.longitude}';
}

void locationFailed(err){
  print('Position Error: ${err.error} ${err.code}');
}