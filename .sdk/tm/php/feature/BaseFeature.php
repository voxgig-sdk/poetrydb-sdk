<?php
declare(strict_types=1);

// Poetrydb SDK base feature

class PoetrydbBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    // Positions this feature when added via the client `extend` option:
    // "__before__" / "__after__" / "__replace__" name an already-added
    // feature (mirrors the ts feature `_options`). Declared so setting it
    // on an extension instance avoids the dynamic-property deprecation.
    public ?array $_options = null;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(PoetrydbContext $ctx, array $options): void {}
    public function PostConstruct(PoetrydbContext $ctx): void {}
    public function PostConstructEntity(PoetrydbContext $ctx): void {}
    public function SetData(PoetrydbContext $ctx): void {}
    public function GetData(PoetrydbContext $ctx): void {}
    public function GetMatch(PoetrydbContext $ctx): void {}
    public function SetMatch(PoetrydbContext $ctx): void {}
    public function PrePoint(PoetrydbContext $ctx): void {}
    public function PreSpec(PoetrydbContext $ctx): void {}
    public function PreRequest(PoetrydbContext $ctx): void {}
    public function PreResponse(PoetrydbContext $ctx): void {}
    public function PreResult(PoetrydbContext $ctx): void {}
    public function PreDone(PoetrydbContext $ctx): void {}
    public function PreUnexpected(PoetrydbContext $ctx): void {}
}
