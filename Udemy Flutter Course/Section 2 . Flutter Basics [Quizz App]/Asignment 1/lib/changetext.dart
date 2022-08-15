import 'package:flutter/material.dart';

class ChangeText extends StatelessWidget {
  final String _greetMessage;

  ChangeText(this._greetMessage);

  @override
  Widget build(BuildContext context) {
    return Container(
      width: double.infinity,
      margin: EdgeInsets.all(10),
      child: Text(
        _greetMessage,
        style: TextStyle(
          fontSize: 20,
        ),
        textAlign: TextAlign.center,
      ),
    );
  }
}
