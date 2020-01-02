import 'dart:convert';
import 'dart:html';
import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/services/secureapi.dart';

class ForgotForm extends FormState {
  EmailInputElement _identity;
  ParagraphElement _error;

  ForgotForm(String idElem, String identityElem, String submitBtn)
      : super(idElem, submitBtn) {
    _identity = querySelector(identityElem);
    _error = querySelector("${idElem}Err");

    querySelector(submitBtn).onClick.listen(onSend);
  }

  String get identity {
    return _identity.value;
  }

  void onSend(Event e) async {
    if (isFormValid()) {
      disableSubmit(true);

      var result = await sendForgot(identity);
      var obj = jsonDecode(result.response);

      if (result.status == 200) {
        final fkey = obj['Data'];
        print(fkey);
      } else {
        _error.text = obj['Error'];
      }
    }
  }
}
