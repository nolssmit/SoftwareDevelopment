import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../providers/product.dart';

class ProductItem extends StatelessWidget {
  final String id;
  final String title;
  final String imageUrl;

  ProductItem({
    required this.id,
    required this.title,
    required this.imageUrl,
  });

  @override
  Widget build(BuildContext context) {
    // return GridTile(
    //   footer: GridTileBar(
    //     backgroundColor: Colors.black87,
    //     leading: IconButton(
    //       icon: const Icon(Icons.favorite, size: 10),
    //       onPressed: () {},
    //     ),
    //     title: SingleChildScrollView(
    //       scrollDirection: Axis.horizontal,
    //       child: Text(
    //         title,
    //         textAlign: TextAlign.center,
    //       ),
    //     ),
    //     trailing: IconButton(
    //       icon: const Icon(Icons.shopping_cart, size: 10),
    //       onPressed: () {},
    //     ),
    //   ),
    //   child: Image.network(
    //     imageUrl,
    //     fit: BoxFit.cover,
    //   ),
    // );


    return ClipRRect(
      borderRadius: BorderRadius.circular(10),
      child: GridTile(
        child: GestureDetector(
          onTap: () {
            // Navigator.of(context).pushNamed(
            //   ProductDetailScreen.routeName,
            //   arguments: product.id,
          //  );
          },
          child: Image.network(
            product.imageUrl,
            fit: BoxFit.cover,
          ),
        ),
        footer: GridTileBar(
          backgroundColor: Colors.black87,
          // leading: Consumer<Product>(
          //   builder: (ctx, product, _) => IconButton(
          //         icon: Icon(
          //           product.isFavorite ? Icons.favorite : Icons.favorite_border,
          //         ),
          //         color: Theme.of(context).accentColor,
          //         onPressed: () {
          //           product.toggleFavoriteStatus();
          //         },
          //       ),
          // ),
          title: Text(
            product.title,
            textAlign: TextAlign.center,
          ),
          trailing: IconButton(
            icon: Icon(
              Icons.shopping_cart,
            ),
            onPressed: () {
//              cart.addItem(product.id, product.price, product.title);
            },
            color: Theme.of(context).accentColor,
          ),
        ),
      ),
    );

  }
}
