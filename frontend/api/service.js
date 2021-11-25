const Services = [
    {
        title: 'Share code',
        icon: 'file_copy',
        description: 'View and share code temp',
        link: '/list-tools/json-encode',
    },
    {
        title: 'JSON Decode',
        icon: 'file_copy',
        description: 'Decode json',
        link: '/list-tools/json-encode',
    }, 
    {
        title: 'JSON Beautify',
        icon: 'file_copy',
        description: 'format json code',
        link: '/list-tools/json-encode',
    }, 
    {
        title: 'Base64 Encode',
        icon: 'file_copy',
        description: 'Convert text to base64',
        link: '/list-tools/base64-encode',
    }, 
    {
        title: 'Base64 Decode',
        icon: 'file_copy',
        description: 'Convert base64 to text',
        link: '/list-tools/base64-decode',
    }, 
    {
        title: 'URL Encode',
        icon: 'file_copy',
        description: 'Encode URL',
        link: '/list-tools/url-encode',
    }, 
    {
        title: 'URL Decode',
        icon: 'file_copy',
        description: 'Decode URL',
        link: '/list-tools/url-decode',
    }, 
    {
        title: 'UTF8 Encode',
        icon: 'file_copy',
        description: 'Template PSD',
        link: '/list-tools/utf8-encode',
    }, 
    {
        title: 'UTF8 Decode',
        icon: 'file_copy',
        description: 'Template PSD',
        link: '/list-tools/utf8-decode',
    }, 
    {
        title: 'JS Format',
        icon: 'file_copy',
        description: 'Format JS',
        link: '/list-tools/js-format',
    },
    {
        title: 'HTML Format',
        icon: 'file_copy',
        description: 'Format HTML',
        link: '/list-tools/html-format',
    },
    {
        title: 'SQL Format',
        icon: 'file_copy',
        description: 'Format SQL',
        link: '/list-tools/sql-format',
    },
    {
        title: 'Mockup Data Generator',
        icon: 'file_copy',
        description: 'Generate data json',
        link: '/list-tools/data-generator',
    },
    {
        title: 'Diff Checker',
        icon: 'file_copy',
        description: 'Check different from 2 string',
        link: '/list-tools/diff-checker',
    },
];

const getServices = (limit) => {
    return (limit) ? Services.slice(0, limit) : Services;
};


export {
    Services,
    getServices
};