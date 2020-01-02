import 'dart:async';
import 'dart:convert';
import 'dart:html';
import 'app.dart';
import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/pathlookup.dart';
import 'package:mango_ui/bodies/register.dart';
import 'package:mango_ui/services/secureapi.dart';

class RegisterForm extends FormState {
  TextInputElement _name;
  EmailInputElement _email;
  PasswordInputElement _password;
  PasswordInputElement _confirm;

  RegisterForm(String idElem, String nameElem, String emailElem,
      String passElem, String confirmElem, String submitBtn)
      : super(idElem, submitBtn) {
    _name = querySelector(nameElem);
    _email = querySelector(emailElem);
    _password = querySelector(passElem);
    _confirm = querySelector(confirmElem);

    querySelector(submitBtn).onClick.listen(onSend);
  }

  String get name {
    return _name.value;
  }

  String get email {
    return _email.value;
  }

  String get password {
    return _password.value;
  }

  String get confirmPassword {
    return _confirm.value;
  }

  void onSend(Event e) {
    if (isFormValid() && passwordsMatch()) {
      disableSubmit(true);
      submitSend().then((obj) {
        disableSubmit(false);
      });
    }
  }

  Future submitSend() async {
    var data = new Register(await getApp(), name, email, password, confirmPassword);
    var result = await sendRegister(data);

    var obj = jsonDecode(result.response);

    print(obj['Data']);
    afterSend(obj['Data']);
  }

  bool passwordsMatch() {
    return password == confirmPassword;
  }

  void afterSend(Object obj) {
    print("We have touchdown {obj}");
  }
}