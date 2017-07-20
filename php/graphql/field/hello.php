<?php
/**
 * User: WindomZ
 * Date: 17-7-20
 */

namespace GraphQLExamples\Field;

use GraphQL\Type\Definition\Type;

function getHelloField()
{
    return [
        'type' => Type::string(),
        'args' => [
            'message' => ['type' => Type::string()],
        ],
        'resolve' => function ($root, $args) {
            return 'Hello ' . $args['message'] . '!';
        }
    ];
}

function getByeField()
{
    return [
        'type' => Type::string(),
//        'resolve' => function ($root, $args) {
//            return $root['bye'];
//        }
    ];
}
