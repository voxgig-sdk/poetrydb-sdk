<?php
declare(strict_types=1);

// Poetrydb SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class PoetrydbFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new PoetrydbBaseFeature();
            case "test":
                return new PoetrydbTestFeature();
            default:
                return new PoetrydbBaseFeature();
        }
    }
}
