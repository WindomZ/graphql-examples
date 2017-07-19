<?php
/**
 * User: WindomZ
 * Date: 17-7-19
 */

require_once __DIR__ . '/../vendor/autoload.php';

use GraphQL\Type\Definition\ObjectType;
use GraphQL\Type\Definition\Type;
use GraphQL\Schema;
use GraphQL\GraphQL;

try {
    $queryType = new ObjectType([
        'name' => 'Hello',
        'fields' => [
            'hello' => [
                'type' => Type::string(),
                'args' => [
                    'message' => ['type' => Type::string()],
                ],
                'resolve' => function ($root, $args) {
                    return $root['prefix'].$args['message'];
                }
            ],
        ],
    ]);

    $mutationType = new ObjectType([
        'name' => 'Calc',
        'fields' => [
            'sum' => [
                'type' => Type::int(),
                'args' => [
                    'x' => ['type' => Type::int()],
                    'y' => ['type' => Type::int()],
                ],
                'resolve' => function ($root, $args) {
                    return $args['x'] + $args['y'];
                },
            ],
        ],
    ]);

    $schema = new Schema([
        'query' => $queryType,
        'mutation' => $mutationType,
    ]);

    // Parse incoming query and variables
    if (isset($_SERVER['CONTENT_TYPE']) && strpos($_SERVER['CONTENT_TYPE'], 'application/json') !== false) {
        $raw = file_get_contents('php://input') ?: '';
        $data = json_decode($raw, true);
    } else {
        $data = $_REQUEST;
    }
    $data += ['query' => null, 'variables' => null];
    $query = $data['query'];
    $variableValues = isset($data['variables']) ? $data['variables'] : null;

    $rootValue = ['prefix' => 'You said: '];
    $result = GraphQL::execute($schema, $query, $rootValue, null, $variableValues);
} catch (\Exception $e) {
    $result = [
        'error' => [
            'message' => $e->getMessage()
        ]
    ];
}
header('Content-Type: application/json; charset=UTF-8');
echo json_encode($result);
