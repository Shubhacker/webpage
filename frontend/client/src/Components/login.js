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
            return front.Login.JwtToken
        }
    }
    // displayVideoTopic(){
    //     var front = this.props.data
    //     if(front.loading){
    //         return (<div>Loading Video...</div>);
    //     }else{
    //         return front.FetchVideo.data.map(book => {
    //             return(
    //                 <li>{book.videoTopic}</li>
    //             );
    //         })
    //     }
    // }
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
