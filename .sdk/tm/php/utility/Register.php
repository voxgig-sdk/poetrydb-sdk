<?php
declare(strict_types=1);

// Poetrydb SDK utility registration

require_once __DIR__ . '/../core/UtilityType.php';
require_once __DIR__ . '/Clean.php';
require_once __DIR__ . '/Done.php';
require_once __DIR__ . '/MakeError.php';
require_once __DIR__ . '/FeatureAdd.php';
require_once __DIR__ . '/FeatureHook.php';
require_once __DIR__ . '/FeatureInit.php';
require_once __DIR__ . '/Fetcher.php';
require_once __DIR__ . '/MakeFetchDef.php';
require_once __DIR__ . '/MakeContext.php';
require_once __DIR__ . '/MakeOptions.php';
require_once __DIR__ . '/MakeRequest.php';
require_once __DIR__ . '/MakeResponse.php';
require_once __DIR__ . '/MakeResult.php';
require_once __DIR__ . '/MakePoint.php';
require_once __DIR__ . '/MakeSpec.php';
require_once __DIR__ . '/MakeUrl.php';
require_once __DIR__ . '/Param.php';
require_once __DIR__ . '/PrepareAuth.php';
require_once __DIR__ . '/PrepareBody.php';
require_once __DIR__ . '/PrepareHeaders.php';
require_once __DIR__ . '/PrepareMethod.php';
require_once __DIR__ . '/PrepareParams.php';
require_once __DIR__ . '/PreparePath.php';
require_once __DIR__ . '/PrepareQuery.php';
require_once __DIR__ . '/ResultBasic.php';
require_once __DIR__ . '/ResultBody.php';
require_once __DIR__ . '/ResultHeaders.php';
require_once __DIR__ . '/TransformRequest.php';
require_once __DIR__ . '/TransformResponse.php';

PoetrydbUtility::setRegistrar(function (PoetrydbUtility $u): void {
    $u->clean = [PoetrydbClean::class, 'call'];
    $u->done = [PoetrydbDone::class, 'call'];
    $u->make_error = [PoetrydbMakeError::class, 'call'];
    $u->feature_add = [PoetrydbFeatureAdd::class, 'call'];
    $u->feature_hook = [PoetrydbFeatureHook::class, 'call'];
    $u->feature_init = [PoetrydbFeatureInit::class, 'call'];
    $u->fetcher = [PoetrydbFetcher::class, 'call'];
    $u->make_fetch_def = [PoetrydbMakeFetchDef::class, 'call'];
    $u->make_context = [PoetrydbMakeContext::class, 'call'];
    $u->make_options = [PoetrydbMakeOptions::class, 'call'];
    $u->make_request = [PoetrydbMakeRequest::class, 'call'];
    $u->make_response = [PoetrydbMakeResponse::class, 'call'];
    $u->make_result = [PoetrydbMakeResult::class, 'call'];
    $u->make_point = [PoetrydbMakePoint::class, 'call'];
    $u->make_spec = [PoetrydbMakeSpec::class, 'call'];
    $u->make_url = [PoetrydbMakeUrl::class, 'call'];
    $u->param = [PoetrydbParam::class, 'call'];
    $u->prepare_auth = [PoetrydbPrepareAuth::class, 'call'];
    $u->prepare_body = [PoetrydbPrepareBody::class, 'call'];
    $u->prepare_headers = [PoetrydbPrepareHeaders::class, 'call'];
    $u->prepare_method = [PoetrydbPrepareMethod::class, 'call'];
    $u->prepare_params = [PoetrydbPrepareParams::class, 'call'];
    $u->prepare_path = [PoetrydbPreparePath::class, 'call'];
    $u->prepare_query = [PoetrydbPrepareQuery::class, 'call'];
    $u->result_basic = [PoetrydbResultBasic::class, 'call'];
    $u->result_body = [PoetrydbResultBody::class, 'call'];
    $u->result_headers = [PoetrydbResultHeaders::class, 'call'];
    $u->transform_request = [PoetrydbTransformRequest::class, 'call'];
    $u->transform_response = [PoetrydbTransformResponse::class, 'call'];
});
