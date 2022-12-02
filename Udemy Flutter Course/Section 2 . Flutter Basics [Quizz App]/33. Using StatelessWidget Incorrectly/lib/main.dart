import 'package:flutter/material.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  var questionIndex = 0;

  // const MyApp({Key? key}) : super(key: key);

  void answerQuestion() {
    questionIndex = questionIndex + 1;
    print(questionIndex);
  }

  @override
  Widget build(BuildContext context) {
    var questions = [
      'What\'s your favorite color?',
      'What\'s your favorite animal?',
    ];
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          title: const Text('My First App'),
        ),
        body: Column(
          children: [
            Text(questions[questionIndex]),
            RaisedButton(
              child: const Text('Answer 1'),
              onPressed: answerQuestion,
            ),
            RaisedButton(
              child: const Text('Answer 2'),
              onPressed: () => print('Button 2 pressed'),
            ),
            RaisedButton(
              child: const Text('Answer 3'),
              onPressed: () {
                print('Button 3 pressed');
              },
            ),
          ],
        ),
      ),
    );
  }
}
