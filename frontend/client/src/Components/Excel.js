import React, {Component} from "react";
import {gql} from 'apollo-boost';
import {graphql} from "react-apollo";
import {render} from "@testing-library/react";
import {Token} from "graphql";

const GetLoginQuery = gql`
query{
    CreateExcelForUser{
      Error
      Message
    }
  }
`

class Excel extends Component{
    displayVideoLink(){
        var front = this.props.data
        if(front.loading){
            return (<div>Loading Video...</div>);
        }else{
            return front.CreateExcelForUser.map(book => {
                return <div>book.Message</div>
            })
        }
    }
    render()
    {
        return (
            <div>
                <h1>{this.displayVideoLink()}</h1>
            </div>
        );
    }
}

export default graphql(GetLoginQuery) (Excel);
