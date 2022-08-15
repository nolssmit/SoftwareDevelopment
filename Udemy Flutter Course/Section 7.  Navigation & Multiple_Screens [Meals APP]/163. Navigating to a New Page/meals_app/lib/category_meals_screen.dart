import 'package:flutter/material.dart';

class CategoryMealsScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('The recipies'),
      ),
      body: Center(
        child: Text(
          'The recipies for that category!',
        ),
      ),
    );
  }
}
