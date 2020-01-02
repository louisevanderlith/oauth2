import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/pathlookup.dart';

Future<HttpRequest> sendResetRequest(
    String resetLink, String username, String email) async {
  final data = jsonEncode({
    "Body": resetLink,
    "Name": username,
    "To": email,
    "Email": "Auth.APP",
    "TemplateName": "resetrequest.html",
  });

  return _sendCommsMessage(data);
}

Future<HttpRequest> sendRegistration(String verifyLink, String username, String email) async {
  final data = jsonEncode({
    "Body": verifyLink,
    "Name": username,
    "To": email,
    "Email": "Auth.APP",
    "TemplateName": "registration.html",
  });

  return _sendCommsMessage(data);
}

Future<HttpRequest> _sendCommsMessage(Object data) async {
  var url = await buildPath("Comms.API", "message", new List<String>());
  
  final compltr = new Completer<HttpRequest>();
  final request = HttpRequest();
  request.open("POST", url);
  request.setRequestHeader("Content-Type", "text/json;charset=UTF-8");
  request.onLoadEnd
      .listen((e) => compltr.complete(request), onError: compltr.completeError);
  request.onError.listen(compltr.completeError);
  request.onProgress.listen(onProgress);
  request.send(data);

  return compltr.future;
}

void onProgress(ProgressEvent e) {
  if (e.lengthComputable) {
    print('Progress... ${e.total}/${e.loaded}');
  }
}
/*
submitSend() async {
  var url = await buildPath("Comms.API", "message", new List<String>());
  var data = jsonEncode({
    "Body": message,
    "Email": email,
    "Name": name,
    "Phone": phone,
    "To": ""
  });

  var resp =
      await HttpRequest.requestCrossOrigin(url, method: "POST", sendData: data);
  var content = jsonDecode(resp);

  if (content['Error'] != "") {
    _error.text = content['Error'];
  } else {
    window.alert(content['Data']);
  }
}
*/