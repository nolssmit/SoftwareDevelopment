import 'package:flutter/material.dart';

class Answer extends StatelessWidget {
  final VoidCallback selectHandler;

  Answer(this.selectHandler);

  @override
  Widget build(BuildContext context) {
    return Container(
      width: double.infinity,
      color: Colors.grey,
      child: RaisedButton(
        child: const Text('Answer 1'),
        onPressed: selectHandler,
      ),
    );
  }
}
