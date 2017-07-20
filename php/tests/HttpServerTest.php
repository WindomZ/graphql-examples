<?php
/**
 * User: WindomZ
 * Date: 17-7-20
 */

namespace GraphQLExamples\Test;

class HttpServerTest extends \PHPUnit_Framework_TestCase
{
    public function testHttpHello()
    {
        self::assertEquals(
            file_get_contents("http://localhost:8080/example?graphql=query{hello(message:\"World\")}"),
            '{"data":{"hello":"Hello World!"}}'
        );

        self::assertEquals(
            file_get_contents("http://localhost:8080/example?graphql=query{bye}"),
            '{"data":{"bye":"Goodbye!"}}'
        );
    }

    public function testHttpCalc()
    {
        self::assertEquals(
            file_get_contents("http://localhost:8080/example?graphql=mutation{sum(x:1,y:2)}"),
            '{"data":{"sum":3}}'
        );
    }

    public function testHttpUser()
    {
        self::assertEquals(
            file_get_contents("http://localhost:8080/example?graphql=query{user(id:\"1\"){name}}"),
            '{"data":{"user":{"name":"Name-1"}}}'
        );

        self::assertEquals(
            file_get_contents("http://localhost:8080/example?graphql=query{user(id:\"id\"){name}}"),
            '{"data":{"user":{"name":"Name"}}}'
        );

        $chinese = urlencode("\"编号\"");
        self::assertEquals(
            file_get_contents("http://localhost:8080/example?graphql=query{user(id:$chinese){name}}"),
            '{"data":{"user":{"name":"名字"}}}'
        );
    }
}