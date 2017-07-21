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
            file_get_contents("http://localhost:8080/graphql?query=query{hello(message:\"World\")}"),
            '{"data":{"hello":"Hello World!"}}'
        );

        self::assertEquals(
            file_get_contents("http://localhost:8080/graphql?query=query{bye}"),
            '{"data":{"bye":"Goodbye!"}}'
        );
    }

    public function testHttpCalc()
    {
        try {
            $url = 'http://localhost:8080/graphql';
            $data = array('query' => 'mutation{sum(x:1,y:2)}');
            $options = array(
                'http' => array(
                    'header' => "Content-type: application/x-www-form-urlencoded\r\n",
                    'method' => 'POST',
                    'content' => http_build_query($data)
                )
            );
            $context = stream_context_create($options);
            $response = file_get_contents($url, false, $context);

            self::assertEquals(
                $response,
                '{"data":{"sum":3}}'
            );
        } catch (\Throwable $t) {
            self::assertEmpty($t);
        } catch (\Exception $e) {
            self::assertEmpty($e);
        }
    }

    public function testHttpUser()
    {
        self::assertEquals(
            file_get_contents("http://localhost:8080/graphql?query=query{user(id:\"1\"){name}}"),
            '{"data":{"user":{"name":"+86-13888888888"}}}'
        );

        self::assertEquals(
            file_get_contents("http://localhost:8080/graphql?query=query{user(id:\"id\"){name}}"),
            '{"data":{"user":{"name":"Name"}}}'
        );

        $chinese = urlencode("\"编号\"");
        self::assertEquals(
            file_get_contents("http://localhost:8080/graphql?query=query{user(id:$chinese){name}}"),
            '{"data":{"user":{"name":"名字"}}}'
        );
    }
}