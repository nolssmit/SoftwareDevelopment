import 'package:flutter/material.dart';

import './catergories_screen.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'DeliMeals',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      // home is the entry point of your app, the root screen
      home: CategoriesScreen(),
    );
  }
}
