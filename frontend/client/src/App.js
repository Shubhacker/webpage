import React, {Component} from "react";
import BookList from "./Components/BookList";
import Login from "./Components/login"
import ApolloClient from "apollo-boost"
import {ApolloProvider} from 'react-apollo'

const client = new ApolloClient({
    uri:'http://localhost:8080/query'

})

class App extends Component{
  render() {
      return (
          <ApolloProvider client={client}>
          <div id="main">
              <h1>Blogging</h1>
              <BookList> </BookList>
          </div>
          <div id="JWT">
          <Login></Login>
          </div>
          </ApolloProvider>
      );
  }

}

export default App;
