import 'package:flutter/material.dart';

class TextControl extends StatelessWidget {
  final Function selectHandler;

  TextControl(this.selectHandler);

  @override
  Widget build(BuildContext context) {
    return Container(
      width: double.infinity,
      child: RaisedButton(
        color: Colors.blue,
        textColor: Colors.white,
        child: const Text('Click to continue...',),
        onPressed: selectHandler,
      ),
    );
  }
}
