extern crate proc_macro;
use crate::proc_macro::TokenStream;
use quote::quote;
use syn;

fn impl_hello_macro(ast: &syn::DeriveInput) -> TokenStream {
    let name = &ast.ident; // get struct name
    let gen = quote! {
        impl HelloMacro for #name {
            fn hello_macro() {
                println!("hello, in my macro, my name is {}", stringify!(#name))
                // 根据需要，stringify函数可以允许定制一个复杂对象的特定属性如何被格式化。 
            }
        }
    };
    gen.into()
}

#[proc_macro_derive(HelloMacro)]
pub fn hello_macro_derive(input: TokenStream) -> TokenStream {
    let ast = syn::parse(input).unwrap();   // parse a struct
    impl_hello_macro(&ast)
}



#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}
