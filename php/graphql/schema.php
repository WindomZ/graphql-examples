<?php
/**
 * User: WindomZ
 * Date: 17-7-20
 */

namespace GraphQLExamples\Schema;

require_once __DIR__ . '/../graphql/field/hello.php';
require_once __DIR__ . '/../graphql/field/calc.php';

use GraphQL\Type\Definition\ObjectType;
use GraphQL\Schema;
use GraphQLExamples\Field;

function getSchema()
{
    return new Schema([
        'query' => new ObjectType([
            'name' => 'Query',
            'fields' => [
                'hello' => Field\getHelloField(),
            ],
        ]),
        'mutation' => new ObjectType([
            'name' => 'Calc',
            'fields' => [
                'sum' => Field\getSumField(),
            ],
        ]),
    ]);
}

