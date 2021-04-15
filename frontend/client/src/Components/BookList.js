import React, {Component} from "react";
import {gql} from 'apollo-boost';
import {graphql} from "react-apollo";
import {render} from "@testing-library/react";

const GetBookQuery = gql`
query{
FetchVideo(input:{

}){
  data{
    video_link
    paid
    videoTopic
    book_name
    tool_name
    status
  }
}  
}
`

class BookList extends Component{
    displayVideoPaid(){
        var front = this.props.data
        if(front.loading){
            return (<div>Loading Video...</div>);
        }else{
            return front.FetchVideo.data.map(book => {
                return (
                    <div>{book.paid}</div>
                )
            })
        }
    }
    displayVideoLink(){
        var front = this.props.data
        if(front.loading){
            return (<div>Loading Video...</div>);
        }else{
            return front.FetchVideo.data.map(book => {
                return(
                    <li>{book.video_link}</li>
            );
            })
        }
    }
    displayVideoTopic(){
        var front = this.props.data
        if(front.loading){
            return (<div>Loading Video...</div>);
        }else{
            return front.FetchVideo.data.map(book => {
                return(
                    <li>{book.videoTopic}</li>
                );
            })
        }
    }
    render()
    {
        return (
            <div>
                <ul id="book-list">
                    {this.displayVideoLink()}
                    {this.displayVideoTopic()}
                    {this.displayVideoPaid()}
                </ul>

            </div>
        );
    }
}

export default graphql(GetBookQuery) (BookList);
