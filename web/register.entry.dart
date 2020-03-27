import 'package:OAUTH2.APP/registerform.dart';

void main() {
  print("Running Register.Entry");

  new RegisterForm("#frmRegister", "#txtName", "#txtEmail", "#txtPassword",
      "#txtConfirmPass", "#btnSubmit");
}
