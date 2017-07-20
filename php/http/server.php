<?php
/**
 * User: WindomZ
 * Date: 17-7-19
 */

require_once __DIR__ . '/../../vendor/autoload.php';
require_once __DIR__ . '/../graphql/schema.php';

use GraphQL\GraphQL;
use GraphQLExamples\Schema;

try {
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
    $result = GraphQL::execute(Schema\getSchema(), $query, $rootValue, null, $variableValues);
} catch (\Exception $e) {
    $result = [
        'error' => [
            'message' => $e->getMessage()
        ]
    ];
}
header('Content-Type: application/json; charset=UTF-8');
echo json_encode($result);
