import 'package:flutter/material.dart';

import 'changetext.dart';
import 'textcontrol.dart';

void main() => runApp(MyApp());

class MyApp extends StatefulWidget {
  @override
  State<StatefulWidget> createState() {
    // TODO: implement createState
    return _MyAppState();
  }
}

class _MyAppState extends State<MyApp> {
  String _greetMessage = 'Hello Mate!';

  void _changeGreetMessage() {
    setState(() {
      _greetMessage = 'How are you today!';
    });
//   print(_newMessage);
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          centerTitle: true,
          title: Text('Assignment 1'),
        ),
        body: Column(
          children: [
            ChangeText(_greetMessage),
            TextControl(_changeGreetMessage),
          ],
        ),
      ),
    );
  }
}
