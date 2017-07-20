<?php
/**
 * User: WindomZ
 * Date: 17-7-20
 */

namespace GraphQLExamples\Schema;

require_once __DIR__ . '/../graphql/field/Hello.php';
require_once __DIR__ . '/../graphql/field/Calc.php';
require_once __DIR__ . '/../graphql/field/User.php';

use GraphQL\Type\Definition\ObjectType;
use GraphQL\Schema;
use GraphQLExamples\Field;

function getSchema()
{
    return new Schema([
        'query' => new ObjectType([
            'name' => 'Query',
            'fields' => [
                'hello' => Field\Hello::getHelloField(),
                'bye' => Field\Hello::getByeField(),
                'user' => Field\User::getUserField(),
            ],
        ]),
        'mutation' => new ObjectType([
            'name' => 'Calc',
            'fields' => [
                'sum' => Field\Calc::getSumField(),
            ],
        ]),
    ]);
}

