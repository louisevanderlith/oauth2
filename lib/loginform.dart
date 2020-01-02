import 'dart:convert';
import 'dart:html';
import 'package:mango_ui/bodies/login.dart';
import 'package:mango_ui/services/secureapi.dart';

import 'package:mango_ui/formstate.dart';
import 'app.dart';

class LoginForm extends FormState {
  EmailInputElement _email;
  PasswordInputElement _password;
  ParagraphElement _error;

  LoginForm(
      String idElem, String emailElem, String passwordElem, String submitBtn)
      : super(idElem, submitBtn) {
    _email = querySelector(emailElem);
    _password = querySelector(passwordElem);
    _error = querySelector("#frmError");

    querySelector(submitBtn).onClick.listen(onSend);
  }

  String get email {
    return _email.value;
  }

  String get password {
    return _password.value;
  }

  void onSend(Event e) async {
    if (isFormValid()) {
      disableSubmit(true);

      final app = await getApp();
      final data = new Login(app, email, password);
      var result = await sendLogin(data);
      var obj = jsonDecode(result.response);

      if (result.status == 200) {
        afterSend(obj['Data']);
      } else {
        _error.text = obj['Error'];
      }
    }
  }
  
  void afterSend(String sessionID) {
    var finalURL = window.localStorage['return'];
    finalURL += "?access_token=" + sessionID;

    window.location.replace(finalURL);
  }
}
