import React, {Component} from "react";
import {gql} from 'apollo-boost';
import {graphql} from "react-apollo";
import {render} from "@testing-library/react";
import {Token} from "graphql";

const GetLoginQuery = gql`
query{
  Login(input:{
    userName:"shubham"
    Password:"Check_123"
  }){
    JwtToken
    Error
  }
}
`

class login extends Component{
    displayVideoLink(){
        var front = this.props.data
        if(front.loading){
            return (<div>Loading Token...</div>);
        }else{
            sessionStorage.setItem("userSessionKey", front.Login.JwtToken);
            // return front.Login.JwtToken
        }
    }
    render()
    {
        return (
            <div>
                <ul id="Login">
                    {this.displayVideoLink()}
                </ul>

            </div>
        );
    }
}

export default graphql(GetLoginQuery) (login);
