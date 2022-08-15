import 'package:flutter/material.dart';

class NewTransaction extends StatelessWidget {
    final titleController = TextEditingController();
  final amountController = TextEditingController();
  //const NewTransaction({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return           Card(
            elevation: 5,
            child: Container(
              padding: EdgeInsets.all(10),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.end,
                children: <Widget>[
                  TextField(
                    decoration: InputDecoration(labelText: 'Title'),
                    controller: titleController,
                    // onChanged: (val) {
                    //   titleInput = val;
                    // },
                  ),
                  TextField(
                    decoration: InputDecoration(labelText: 'Amount'),
                    controller: amountController,
                    // onChanged: (val) => amountInput = val,
                  ),
                  FlatButton(
                    child: Text('Add Transaction'),
                    onPressed: () {
                      print(titleController.text);
                      print(amountController.text);
                    },
                    textColor: Colors.purple,
                  ),
                ],
              ),
            ),
          );
  }
}
